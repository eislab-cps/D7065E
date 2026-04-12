package server

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/graph"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/handlers"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type Server struct {
	Port     int
	EditMode bool
	store    *store.MemoryStore
	hub      *websocket.Hub
	dataFS   embed.FS
	webFS    embed.FS
}

func New(port int, dataFS embed.FS, webFS embed.FS, editMode bool) *Server {
	st := store.NewMemoryStore()
	hub := websocket.NewHub(st)

	return &Server{
		Port:     port,
		EditMode: editMode,
		store:    st,
		hub:      hub,
		dataFS:   dataFS,
		webFS:    webFS,
	}
}

func (s *Server) loadBuildingData() error {
	building := model.Building{
		Name: "A-Building (LTU)",
		Levels: []model.Level{
			{ID: "level0", Label: "Floor 0"},
			{ID: "level1", Label: "Floor 1"},
			{ID: "level2", Label: "Floor 2"},
		},
	}
	s.store.SetBuilding(building)

	// Load floor data
	for _, level := range building.Levels {
		path := fmt.Sprintf("data/abuilding/%s/floorplan_data.json", level.ID)
		data, err := s.dataFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}
		var floorData model.FloorData
		if err := json.Unmarshal(data, &floorData); err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}
		s.store.SetFloorData(level.ID, &floorData)
		log.Printf("Loaded %s: %d rooms, %d walls", level.ID, len(floorData.Rooms), len(floorData.Walls))
	}

	// Load cross-floor edges
	cfePath := "data/abuilding/cross_floor_edges.json"
	cfeData, err := s.dataFS.ReadFile(cfePath)
	if err != nil {
		log.Printf("No cross-floor edges: %v", err)
	} else {
		var edges []model.CrossFloorEdge
		if err := json.Unmarshal(cfeData, &edges); err != nil {
			log.Printf("Failed to parse cross-floor edges: %v", err)
		} else {
			s.store.SetCrossFloorEdges(edges)
			log.Printf("Loaded %d cross-floor edges", len(edges))
		}
	}

	// Build multi-floor walkable graph
	levelOrder := []string{}
	for _, level := range building.Levels {
		levelOrder = append(levelOrder, level.ID)
	}
	mfg := graph.BuildMultiFloorGraph(s.store.GetFloors(), levelOrder, s.store.GetCrossFloorEdges())
	s.store.SetMultiFloorGraph(mfg)
	log.Printf("Built multi-floor walkable graph: %d nodes, %d edges", len(mfg.Nodes), len(mfg.Edges))

	return nil
}

func (s *Server) Start() error {
	// Set up log file
	logFile, err := os.OpenFile("buildsim.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = multiWriter

	if err := s.loadBuildingData(); err != nil {
		return fmt.Errorf("failed to load building data: %w", err)
	}

	// Start session purger (check every 5 min, purge after 1 hour)
	s.hub.StartPurger(5*time.Minute, 1*time.Hour)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		// Config endpoint
		editMode := s.EditMode
		api.GET("/config", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"edit_mode": editMode})
		})

		bh := &handlers.BuildingHandlers{Store: s.store}
		api.GET("/building", bh.GetBuilding)
		api.GET("/building/floors/:level", bh.GetFloor)
		api.GET("/building/cross-floor-edges", bh.GetCrossFloorEdges)

		gh := &handlers.GraphHandlers{Store: s.store}
		api.GET("/graph", gh.GetGraph)
		api.GET("/graph/route", gh.GetRoute)

		eh := &handlers.EquipmentHandlers{Store: s.store, Hub: s.hub}
		api.POST("/equipment", eh.Create)
		api.POST("/equipment/bulk", eh.BulkCreate)
		api.GET("/equipment", eh.List)
		api.GET("/equipment/:id", eh.Get)
		api.PUT("/equipment/:id", eh.Update)
		api.DELETE("/equipment/:id", eh.Delete)
		api.POST("/equipment/notify", eh.Notify)

		senh := &handlers.SensorHandlers{Store: s.store}
		api.POST("/equipment/:id/sensors", senh.AddSensor)
		api.GET("/equipment/:id/sensors", senh.ListSensors)
		api.DELETE("/sensors/:id", senh.DeleteSensor)
		api.PUT("/sensors/:id/value", senh.SetValue)

		acth := &handlers.ActuatorHandlers{Store: s.store}
		api.POST("/equipment/:id/actuators", acth.AddActuator)
		api.GET("/equipment/:id/actuators", acth.ListActuators)
		api.DELETE("/actuators/:id", acth.DeleteActuator)
		api.PUT("/actuators/:id/state", acth.SetState)

		sh := &handlers.SessionHandlers{Store: s.store, Hub: s.hub}
		api.GET("/sessions", sh.List)
		api.POST("/sessions", sh.Create)
		api.GET("/sessions/:id", sh.Get)
		api.DELETE("/sessions/:id", sh.Delete)
		api.PUT("/sessions/:id/viewport", sh.SetViewport)
		api.PUT("/sessions/:id/highlights", sh.SetHighlights)
		api.PUT("/sessions/:id/occupancy", sh.SetOccupancy)
		api.PUT("/sessions/:id/route", sh.SetRoute)
		api.PUT("/sessions/:id/coverage", sh.SetCoverage)

		// Icons
		api.GET("/icons/:name", func(c *gin.Context) {
			name := c.Param("name")
			data, err := s.dataFS.ReadFile("data/equipment/icons/" + name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "icon not found"})
				return
			}
			c.Data(http.StatusOK, "image/svg+xml", data)
		})
	}

	// WebSocket
	sh := &handlers.SessionHandlers{Store: s.store, Hub: s.hub}
	r.GET("/ws/:id", sh.HandleWebSocket)

	// Serve static web files
	webSub, err := fs.Sub(s.webFS, "web")
	if err != nil {
		return fmt.Errorf("failed to create web sub-filesystem: %w", err)
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.FS(webSub))))

	addr := fmt.Sprintf("0.0.0.0:%d", s.Port)
	log.Printf("BuildSim server starting on http://%s", addr)
	return r.Run(addr)
}
