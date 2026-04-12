package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/handlers"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
	"github.com/eislab-cps/buildingsim/pkg/graph"
)

func setupTestRouter() (*gin.Engine, *store.MemoryStore) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	st := store.NewMemoryStore()
	hub := websocket.NewHub(st)

	// Load minimal test data
	st.SetBuilding(model.Building{
		Name: "Test Building",
		Levels: []model.Level{
			{ID: "level0", Label: "Floor 0"},
		},
	})
	st.SetFloorData("level0", &model.FloorData{
		Page: model.Page{Width: 100, Height: 100},
		Rooms: []model.Room{
			{ID: 0, Name: "R001", Area: 50, Center: [2]float64{25, 25}, Type: "corridor"},
			{ID: 1, Name: "R002", Area: 30, Center: [2]float64{50, 50}, Type: "room"},
			{ID: 2, Name: "R003", Area: 30, Center: [2]float64{75, 75}, Type: "room"},
		},
		Graph: &model.NavGraph{
			Nodes: []model.NavNode{
				{ID: 0, Name: "R001", X: 25, Y: 25},
				{ID: 1, Name: "R002", X: 50, Y: 50},
				{ID: 2, Name: "R003", X: 75, Y: 75},
			},
			Edges: []model.NavEdge{
				{From: 0, To: 1, Weight: 35.4},
				{From: 1, To: 2, Weight: 35.4},
			},
		},
		WalkableGraph: &model.NavGraph{
			Nodes: []model.NavNode{
				{ID: 0, Name: "R001", X: 25, Y: 25, Type: "corridor"},
				{ID: 1, Name: "R002", X: 50, Y: 50, Type: "room"},
				{ID: 2, Name: "R003", X: 75, Y: 75, Type: "room"},
			},
			Edges: []model.NavEdge{
				{From: 0, To: 1, Weight: 35.4},
				{From: 1, To: 2, Weight: 35.4},
			},
		},
	})

	// Build multi-floor graph
	mfg := graph.BuildMultiFloorGraph(st.GetFloors(), []string{"level0"}, nil)
	st.SetMultiFloorGraph(mfg)

	api := r.Group("/api")
	{
		bh := &handlers.BuildingHandlers{Store: st}
		api.GET("/building", bh.GetBuilding)
		api.GET("/building/floors/:level", bh.GetFloor)
		api.GET("/building/cross-floor-edges", bh.GetCrossFloorEdges)

		gh := &handlers.GraphHandlers{Store: st}
		api.GET("/graph", gh.GetGraph)
		api.GET("/graph/route", gh.GetRoute)

		eh := &handlers.EquipmentHandlers{Store: st, Hub: hub}
		api.POST("/equipment", eh.Create)
		api.POST("/equipment/bulk", eh.BulkCreate)
		api.GET("/equipment", eh.List)
		api.GET("/equipment/:id", eh.Get)
		api.PUT("/equipment/:id", eh.Update)
		api.DELETE("/equipment/:id", eh.Delete)
		api.POST("/equipment/notify", eh.Notify)

		senh := &handlers.SensorHandlers{Store: st}
		api.POST("/equipment/:id/sensors", senh.AddSensor)
		api.GET("/equipment/:id/sensors", senh.ListSensors)
		api.DELETE("/sensors/:id", senh.DeleteSensor)
		api.PUT("/sensors/:id/value", senh.SetValue)

		acth := &handlers.ActuatorHandlers{Store: st}
		api.POST("/equipment/:id/actuators", acth.AddActuator)
		api.GET("/equipment/:id/actuators", acth.ListActuators)
		api.DELETE("/actuators/:id", acth.DeleteActuator)
		api.PUT("/actuators/:id/state", acth.SetState)

		oh := &handlers.OccupancyHandlers{Store: st, Hub: hub}
		api.GET("/occupancy", oh.Get)
		api.PUT("/occupancy", oh.Set)

		covh := &handlers.CoverageHandlers{Store: st, Hub: hub}
		api.GET("/coverage", covh.Get)
		api.PUT("/coverage", covh.Set)

		sh := &handlers.SessionHandlers{Store: st, Hub: hub}
		api.GET("/sessions", sh.List)
		api.POST("/sessions", sh.Create)
		api.GET("/sessions/:id", sh.Get)
		api.DELETE("/sessions/:id", sh.Delete)
		api.PUT("/sessions/:id/viewport", sh.SetViewport)
		api.PUT("/sessions/:id/highlights", sh.SetHighlights)
		api.PUT("/sessions/:id/occupancy", sh.SetOccupancy)
		api.PUT("/sessions/:id/route", sh.SetRoute)
		api.PUT("/sessions/:id/coverage", sh.SetCoverage)
	}

	return r, st
}

func doRequest(r *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		reqBody = bytes.NewReader(b)
	}
	req, _ := http.NewRequest(method, path, reqBody)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func parseJSON(w *httptest.ResponseRecorder, v interface{}) {
	json.Unmarshal(w.Body.Bytes(), v)
}

// === Building API Tests ===

func TestGetBuilding(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/building", nil)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var b model.Building
	parseJSON(w, &b)
	if b.Name != "Test Building" {
		t.Fatalf("expected 'Test Building', got '%s'", b.Name)
	}
	if len(b.Levels) != 1 {
		t.Fatalf("expected 1 level, got %d", len(b.Levels))
	}
}

func TestGetFloor(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/building/floors/level0", nil)

	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var fd model.FloorData
	parseJSON(w, &fd)
	if len(fd.Rooms) != 3 {
		t.Fatalf("expected 3 rooms, got %d", len(fd.Rooms))
	}
}

func TestGetFloorNotFound(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/building/floors/level99", nil)

	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// === Equipment API Tests ===

func TestEquipmentCRUD(t *testing.T) {
	r, _ := setupTestRouter()

	// Create
	eq := model.Equipment{
		ID: "temp-1", Name: "Temp Sensor", Type: "temperature_sensor",
		Category: "monitoring", Level: "level0", Room: "R002", Status: "running",
	}
	w := doRequest(r, "POST", "/api/equipment", eq)
	if w.Code != 201 {
		t.Fatalf("create: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// Get
	w = doRequest(r, "GET", "/api/equipment/temp-1", nil)
	if w.Code != 200 {
		t.Fatalf("get: expected 200, got %d", w.Code)
	}
	var got model.Equipment
	parseJSON(w, &got)
	if got.Room != "R002" {
		t.Fatalf("expected room R002, got %s", got.Room)
	}

	// List
	w = doRequest(r, "GET", "/api/equipment", nil)
	if w.Code != 200 {
		t.Fatalf("list: expected 200, got %d", w.Code)
	}
	var list []*model.Equipment
	parseJSON(w, &list)
	if len(list) != 1 {
		t.Fatalf("expected 1 equipment, got %d", len(list))
	}

	// List with filter
	w = doRequest(r, "GET", "/api/equipment?room=R002", nil)
	var filtered []*model.Equipment
	parseJSON(w, &filtered)
	if len(filtered) != 1 {
		t.Fatalf("expected 1 filtered, got %d", len(filtered))
	}
	w = doRequest(r, "GET", "/api/equipment?room=R999", nil)
	parseJSON(w, &filtered)
	if len(filtered) != 0 {
		t.Fatalf("expected 0 filtered, got %d", len(filtered))
	}

	// Update
	eq.Status = "warning"
	w = doRequest(r, "PUT", "/api/equipment/temp-1", eq)
	if w.Code != 200 {
		t.Fatalf("update: expected 200, got %d", w.Code)
	}

	// Delete
	w = doRequest(r, "DELETE", "/api/equipment/temp-1", nil)
	if w.Code != 200 {
		t.Fatalf("delete: expected 200, got %d", w.Code)
	}
	w = doRequest(r, "GET", "/api/equipment/temp-1", nil)
	if w.Code != 404 {
		t.Fatalf("after delete: expected 404, got %d", w.Code)
	}
}

func TestEquipmentDuplicate(t *testing.T) {
	r, _ := setupTestRouter()
	eq := model.Equipment{ID: "dup-1", Name: "Test", Type: "test", Level: "level0", Room: "R001"}
	doRequest(r, "POST", "/api/equipment", eq)
	w := doRequest(r, "POST", "/api/equipment", eq)
	if w.Code != 409 {
		t.Fatalf("expected 409 conflict, got %d", w.Code)
	}
}

// === Sensor API Tests ===

func TestSensorCRUD(t *testing.T) {
	r, _ := setupTestRouter()

	// Create equipment first
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-1", Name: "Test", Type: "test", Level: "level0", Room: "R001",
	})

	// Add sensor
	sen := model.Sensor{ID: "sen-1", Name: "Temperature", Type: "temperature", DataType: "text", Unit: "°C", Value: "21.0"}
	w := doRequest(r, "POST", "/api/equipment/eq-1/sensors", sen)
	if w.Code != 201 {
		t.Fatalf("add sensor: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// List sensors
	w = doRequest(r, "GET", "/api/equipment/eq-1/sensors", nil)
	if w.Code != 200 {
		t.Fatalf("list sensors: expected 200, got %d", w.Code)
	}
	var sensors []model.Sensor
	parseJSON(w, &sensors)
	if len(sensors) != 1 {
		t.Fatalf("expected 1 sensor, got %d", len(sensors))
	}

	// Set value (text)
	w = doRequest(r, "PUT", "/api/sensors/sen-1/value", model.SensorValue{DataType: "text", Value: "23.5"})
	if w.Code != 200 {
		t.Fatalf("set value: expected 200, got %d", w.Code)
	}

	// Verify value changed
	w = doRequest(r, "GET", "/api/equipment/eq-1", nil)
	var eq model.Equipment
	parseJSON(w, &eq)
	if eq.Sensors[0].Value != "23.5" {
		t.Fatalf("expected value 23.5, got %s", eq.Sensors[0].Value)
	}

	// Set value (binary)
	boolVal := true
	w = doRequest(r, "PUT", "/api/sensors/sen-1/value", model.SensorValue{DataType: "binary", BinaryValue: &boolVal})
	if w.Code != 200 {
		t.Fatalf("set binary value: expected 200, got %d", w.Code)
	}

	// Delete sensor
	w = doRequest(r, "DELETE", "/api/sensors/sen-1", nil)
	if w.Code != 200 {
		t.Fatalf("delete sensor: expected 200, got %d", w.Code)
	}
}

// === Actuator API Tests ===

func TestActuatorCRUD(t *testing.T) {
	r, _ := setupTestRouter()

	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-2", Name: "Door", Type: "door_lock", Level: "level0", Room: "R001",
	})

	// Add actuator
	act := model.Actuator{ID: "act-1", Name: "Lock", Type: "lock_control", State: "locked"}
	w := doRequest(r, "POST", "/api/equipment/eq-2/actuators", act)
	if w.Code != 201 {
		t.Fatalf("add actuator: expected 201, got %d: %s", w.Code, w.Body.String())
	}

	// List actuators
	w = doRequest(r, "GET", "/api/equipment/eq-2/actuators", nil)
	var actuators []model.Actuator
	parseJSON(w, &actuators)
	if len(actuators) != 1 {
		t.Fatalf("expected 1 actuator, got %d", len(actuators))
	}

	// Set state
	w = doRequest(r, "PUT", "/api/actuators/act-1/state", model.ActuatorState{State: "unlocked"})
	if w.Code != 200 {
		t.Fatalf("set state: expected 200, got %d", w.Code)
	}

	// Verify
	w = doRequest(r, "GET", "/api/equipment/eq-2", nil)
	var eq model.Equipment
	parseJSON(w, &eq)
	if eq.Actuators[0].State != "unlocked" {
		t.Fatalf("expected unlocked, got %s", eq.Actuators[0].State)
	}

	// Delete
	w = doRequest(r, "DELETE", "/api/actuators/act-1", nil)
	if w.Code != 200 {
		t.Fatalf("delete actuator: expected 200, got %d", w.Code)
	}
}

// === Session API Tests ===

func TestSessionLifecycle(t *testing.T) {
	r, _ := setupTestRouter()

	// Create
	w := doRequest(r, "POST", "/api/sessions", nil)
	if w.Code != 201 {
		t.Fatalf("create session: expected 201, got %d", w.Code)
	}
	var sess model.Session
	parseJSON(w, &sess)
	if sess.ID == "" {
		t.Fatal("session ID should not be empty")
	}

	// List
	w = doRequest(r, "GET", "/api/sessions", nil)
	var sessions []*model.Session
	parseJSON(w, &sessions)
	if len(sessions) != 1 {
		t.Fatalf("expected 1 session, got %d", len(sessions))
	}

	// Get
	w = doRequest(r, "GET", "/api/sessions/"+sess.ID, nil)
	if w.Code != 200 {
		t.Fatalf("get session: expected 200, got %d", w.Code)
	}

	// Set viewport
	w = doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/viewport",
		model.Viewport{Room: "R002", Zoom: 2.0, Mode: "3d"})
	if w.Code != 200 {
		t.Fatalf("set viewport: expected 200, got %d", w.Code)
	}

	// Set highlights
	w = doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/highlights",
		[]model.RoomHighlight{{RoomID: 1, Color: "#ff0000", Opacity: 0.8}})
	if w.Code != 200 {
		t.Fatalf("set highlights: expected 200, got %d", w.Code)
	}

	// Set occupancy
	occ := map[int]model.RoomOccupancy{
		1: {Persons: []model.Person{{ID: "p1", Name: "Alice"}}, Aliens: []model.Alien{}},
	}
	w = doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/occupancy", occ)
	if w.Code != 200 {
		t.Fatalf("set occupancy: expected 200, got %d", w.Code)
	}

	// Set coverage
	w = doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/coverage",
		[]model.CoverageZone{{ID: "c1", Name: "WiFi", Room: "R001", Radius: 30, Color: "#00aaff", Opacity: 0.1}})
	if w.Code != 200 {
		t.Fatalf("set coverage: expected 200, got %d", w.Code)
	}

	// Verify session state
	w = doRequest(r, "GET", "/api/sessions/"+sess.ID, nil)
	parseJSON(w, &sess)
	if sess.Viewport.Room != "R002" {
		t.Fatalf("expected viewport room R002, got %s", sess.Viewport.Room)
	}
	if len(sess.Highlights) != 1 {
		t.Fatalf("expected 1 highlight, got %d", len(sess.Highlights))
	}
	if len(sess.Coverage) != 1 {
		t.Fatalf("expected 1 coverage zone, got %d", len(sess.Coverage))
	}

	// Delete
	w = doRequest(r, "DELETE", "/api/sessions/"+sess.ID, nil)
	if w.Code != 200 {
		t.Fatalf("delete session: expected 200, got %d", w.Code)
	}
	w = doRequest(r, "GET", "/api/sessions/"+sess.ID, nil)
	if w.Code != 404 {
		t.Fatalf("after delete: expected 404, got %d", w.Code)
	}
}

// === Graph API Tests ===

func TestGetGraph(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "GET", "/api/graph?level=level0", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var g model.NavGraph
	parseJSON(w, &g)
	if len(g.Nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(g.Nodes))
	}
	if len(g.Edges) != 2 {
		t.Fatalf("expected 2 edges, got %d", len(g.Edges))
	}
}

func TestGetWalkableGraph(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "GET", "/api/graph?level=level0&type=walkable", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var g model.NavGraph
	parseJSON(w, &g)
	if len(g.Nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(g.Nodes))
	}
}

func TestRouteByID(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "GET", "/api/graph/route?from=0&to=2&level=level0", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	var result model.RouteResult
	parseJSON(w, &result)
	if len(result.Path) != 3 {
		t.Fatalf("expected 3 path nodes, got %d", len(result.Path))
	}
	if result.Distance < 1 {
		t.Fatalf("expected positive distance, got %f", result.Distance)
	}
}

func TestRouteByName(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "GET", "/api/graph/route?from_name=R001&to_name=R003&level=level0&type=walkable", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	var result model.RouteResult
	parseJSON(w, &result)
	if len(result.Path) < 2 {
		t.Fatalf("expected at least 2 path nodes, got %d", len(result.Path))
	}
}

func TestRouteNotFound(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "GET", "/api/graph/route?from_name=R001&to_name=NONEXISTENT&level=level0", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// === Equipment Notify Tests ===

func TestBulkCreateEquipment(t *testing.T) {
	r, _ := setupTestRouter()

	items := []model.Equipment{
		{ID: "bulk-1", Name: "A", Type: "test", Category: "monitoring", Level: "level0", Room: "R001"},
		{ID: "bulk-2", Name: "B", Type: "test", Category: "monitoring", Level: "level0", Room: "R002",
			Sensors: []model.Sensor{{ID: "bulk-2-s", Name: "S", Type: "t", DataType: "text", Value: "1"}},
			Actuators: []model.Actuator{{ID: "bulk-2-a", Name: "A", Type: "t", State: "on"}},
		},
		{ID: "bulk-3", Name: "C", Type: "test", Category: "monitoring", Level: "level0", Room: "R003"},
	}
	w := doRequest(r, "POST", "/api/equipment/bulk", items)
	if w.Code != 201 {
		t.Fatalf("expected 201, got %d: %s", w.Code, w.Body.String())
	}
	var resp map[string]interface{}
	parseJSON(w, &resp)
	if int(resp["created"].(float64)) != 3 {
		t.Fatalf("expected 3 created, got %v", resp["created"])
	}

	// Verify sensors and actuators were created
	w = doRequest(r, "GET", "/api/equipment/bulk-2", nil)
	var eq model.Equipment
	parseJSON(w, &eq)
	if len(eq.Sensors) != 1 {
		t.Fatalf("expected 1 sensor, got %d", len(eq.Sensors))
	}
	if len(eq.Actuators) != 1 {
		t.Fatalf("expected 1 actuator, got %d", len(eq.Actuators))
	}

	// Duplicates should be skipped
	w = doRequest(r, "POST", "/api/equipment/bulk", items)
	parseJSON(w, &resp)
	if int(resp["created"].(float64)) != 0 {
		t.Fatalf("expected 0 created (duplicates), got %v", resp["created"])
	}
}

func TestDuplicateSensor(t *testing.T) {
	r, _ := setupTestRouter()
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-dup-s", Name: "T", Type: "t", Level: "level0", Room: "R001",
	})
	sen := model.Sensor{ID: "s-dup", Name: "S", Type: "t", DataType: "text"}
	w := doRequest(r, "POST", "/api/equipment/eq-dup-s/sensors", sen)
	if w.Code != 201 {
		t.Fatalf("first add: expected 201, got %d", w.Code)
	}
	w = doRequest(r, "POST", "/api/equipment/eq-dup-s/sensors", sen)
	if w.Code != 409 {
		t.Fatalf("duplicate add: expected 409, got %d: %s", w.Code, w.Body.String())
	}
}

func TestDuplicateActuator(t *testing.T) {
	r, _ := setupTestRouter()
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-dup-a", Name: "T", Type: "t", Level: "level0", Room: "R001",
	})
	act := model.Actuator{ID: "a-dup", Name: "A", Type: "t", State: "off"}
	w := doRequest(r, "POST", "/api/equipment/eq-dup-a/actuators", act)
	if w.Code != 201 {
		t.Fatalf("first add: expected 201, got %d", w.Code)
	}
	w = doRequest(r, "POST", "/api/equipment/eq-dup-a/actuators", act)
	if w.Code != 409 {
		t.Fatalf("duplicate add: expected 409, got %d: %s", w.Code, w.Body.String())
	}
}

func TestGlobalOccupancy(t *testing.T) {
	r, _ := setupTestRouter()

	// Initially empty
	w := doRequest(r, "GET", "/api/occupancy", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Set occupancy
	occ := map[int]model.RoomOccupancy{
		1: {Persons: []model.Person{{ID: "p1", Name: "Alice"}}, Aliens: []model.Alien{{ID: "x1"}}},
	}
	w = doRequest(r, "PUT", "/api/occupancy", occ)
	if w.Code != 200 {
		t.Fatalf("set: expected 200, got %d", w.Code)
	}

	// Verify persists
	w = doRequest(r, "GET", "/api/occupancy", nil)
	var got map[int]model.RoomOccupancy
	parseJSON(w, &got)
	if len(got) != 1 {
		t.Fatalf("expected 1 room, got %d", len(got))
	}
	if len(got[1].Persons) != 1 || got[1].Persons[0].Name != "Alice" {
		t.Fatalf("expected Alice, got %v", got[1])
	}
	if len(got[1].Aliens) != 1 {
		t.Fatalf("expected 1 alien, got %d", len(got[1].Aliens))
	}

}

func TestGlobalCoverage(t *testing.T) {
	r, _ := setupTestRouter()

	// Initially empty
	w := doRequest(r, "GET", "/api/coverage", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Set coverage
	zones := []model.CoverageZone{
		{ID: "z1", Name: "WiFi", Room: "R001", Radius: 30, Color: "#00aaff", Opacity: 0.1},
	}
	w = doRequest(r, "PUT", "/api/coverage", zones)
	if w.Code != 200 {
		t.Fatalf("set: expected 200, got %d", w.Code)
	}

	// Verify persists
	w = doRequest(r, "GET", "/api/coverage", nil)
	var got []model.CoverageZone
	parseJSON(w, &got)
	if len(got) != 1 {
		t.Fatalf("expected 1 zone, got %d", len(got))
	}
	if got[0].Name != "WiFi" {
		t.Fatalf("expected WiFi, got %s", got[0].Name)
	}

	// Clear
	w = doRequest(r, "PUT", "/api/coverage", []model.CoverageZone{})
	if w.Code != 200 {
		t.Fatalf("clear: expected 200, got %d", w.Code)
	}
	w = doRequest(r, "GET", "/api/coverage", nil)
	parseJSON(w, &got)
	if len(got) != 0 {
		t.Fatalf("expected empty after clear, got %d", len(got))
	}
}

func TestEquipmentNotify(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "POST", "/api/equipment/notify", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	parseJSON(w, &resp)
	if resp["version"] == nil {
		t.Fatal("expected version in response")
	}
}

// === Edge Cases ===

func TestEquipmentNotFound(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/equipment/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestSensorNotFound(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "PUT", "/api/sensors/nonexistent/value",
		model.SensorValue{DataType: "text", Value: "1"})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestActuatorNotFound(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "PUT", "/api/actuators/nonexistent/state",
		model.ActuatorState{State: "on"})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestSessionNotFound(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/sessions/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// === Error Handling & Edge Cases ===

func TestCreateEquipmentMissingID(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "POST", "/api/equipment", model.Equipment{
		Name: "No ID", Type: "test", Level: "level0", Room: "R001",
	})
	if w.Code != 400 {
		t.Fatalf("expected 400 for missing ID, got %d", w.Code)
	}
}

func TestCreateSensorMissingID(t *testing.T) {
	r, _ := setupTestRouter()
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-noid", Name: "T", Type: "t", Level: "level0", Room: "R001",
	})
	w := doRequest(r, "POST", "/api/equipment/eq-noid/sensors",
		model.Sensor{Name: "S", Type: "t", DataType: "text"})
	if w.Code != 400 {
		t.Fatalf("expected 400 for missing sensor ID, got %d", w.Code)
	}
}

func TestCreateActuatorMissingID(t *testing.T) {
	r, _ := setupTestRouter()
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-noid2", Name: "T", Type: "t", Level: "level0", Room: "R001",
	})
	w := doRequest(r, "POST", "/api/equipment/eq-noid2/actuators",
		model.Actuator{Name: "A", Type: "t", State: "off"})
	if w.Code != 400 {
		t.Fatalf("expected 400 for missing actuator ID, got %d", w.Code)
	}
}

func TestAddSensorToNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "POST", "/api/equipment/nonexistent/sensors",
		model.Sensor{ID: "s1", Name: "S", Type: "t", DataType: "text"})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestAddActuatorToNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "POST", "/api/equipment/nonexistent/actuators",
		model.Actuator{ID: "a1", Name: "A", Type: "t", State: "off"})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestUpdateNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "PUT", "/api/equipment/nonexistent", model.Equipment{
		Name: "X", Type: "t", Level: "level0", Room: "R001",
	})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "DELETE", "/api/equipment/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteNonexistentSensor(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "DELETE", "/api/sensors/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteNonexistentActuator(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "DELETE", "/api/actuators/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestDeleteNonexistentSession(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "DELETE", "/api/sessions/nonexistent", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestSetViewportOnNonexistentSession(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "PUT", "/api/sessions/nonexistent/viewport",
		model.Viewport{Room: "R001", Zoom: 1.0, Mode: "3d"})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestSetHighlightsOnNonexistentSession(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "PUT", "/api/sessions/nonexistent/highlights",
		[]model.RoomHighlight{{RoomID: 1, Color: "#ff0000", Opacity: 0.5}})
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestInvalidJSON(t *testing.T) {
	r, _ := setupTestRouter()
	req, _ := http.NewRequest("POST", "/api/equipment", bytes.NewReader([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Fatalf("expected 400 for invalid JSON, got %d", w.Code)
	}
}

func TestRouteInvalidFromID(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/graph/route?from=abc&to=2&level=level0", nil)
	if w.Code != 400 {
		t.Fatalf("expected 400 for invalid from ID, got %d", w.Code)
	}
}

func TestRouteInvalidToID(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/graph/route?from=0&to=abc&level=level0", nil)
	if w.Code != 400 {
		t.Fatalf("expected 400 for invalid to ID, got %d", w.Code)
	}
}

func TestGraphNonexistentLevel(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/graph?level=level99", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestListSensorsOnNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/equipment/nonexistent/sensors", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestListActuatorsOnNonexistentEquipment(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/equipment/nonexistent/actuators", nil)
	if w.Code != 404 {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestEmptyEquipmentList(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/equipment", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var list []*model.Equipment
	parseJSON(w, &list)
	if len(list) != 0 {
		t.Fatalf("expected empty list, got %d", len(list))
	}
}

func TestEmptySessionList(t *testing.T) {
	r, _ := setupTestRouter()
	w := doRequest(r, "GET", "/api/sessions", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var list []*model.Session
	parseJSON(w, &list)
	if len(list) != 0 {
		t.Fatalf("expected empty list, got %d", len(list))
	}
}

func TestMultipleSessionsIndependent(t *testing.T) {
	r, _ := setupTestRouter()

	// Create two sessions
	w1 := doRequest(r, "POST", "/api/sessions", nil)
	var s1 model.Session
	parseJSON(w1, &s1)

	w2 := doRequest(r, "POST", "/api/sessions", nil)
	var s2 model.Session
	parseJSON(w2, &s2)

	if s1.ID == s2.ID {
		t.Fatal("sessions should have different IDs")
	}

	// Set highlights on session 1 only
	doRequest(r, "PUT", "/api/sessions/"+s1.ID+"/highlights",
		[]model.RoomHighlight{{RoomID: 1, Color: "#ff0000", Opacity: 0.8}})

	// Session 2 should not have highlights
	w := doRequest(r, "GET", "/api/sessions/"+s2.ID, nil)
	parseJSON(w, &s2)
	if len(s2.Highlights) != 0 {
		t.Fatalf("session 2 should have no highlights, got %d", len(s2.Highlights))
	}

	// Delete session 1, session 2 should still exist
	doRequest(r, "DELETE", "/api/sessions/"+s1.ID, nil)
	w = doRequest(r, "GET", "/api/sessions/"+s2.ID, nil)
	if w.Code != 200 {
		t.Fatalf("session 2 should still exist, got %d", w.Code)
	}
}

func TestConcurrentSensorUpdates(t *testing.T) {
	r, _ := setupTestRouter()

	// Create equipment with sensor
	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-race", Name: "Race Test", Type: "test", Level: "level0", Room: "R001",
	})
	doRequest(r, "POST", "/api/equipment/eq-race/sensors",
		model.Sensor{ID: "s-race", Name: "S", Type: "t", DataType: "text", Value: "0"})

	// Hammer it from 10 goroutines
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			for j := 0; j < 50; j++ {
				doRequest(r, "PUT", "/api/sensors/s-race/value",
					model.SensorValue{DataType: "text", Value: string(rune('0' + n))})
			}
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	// Should still work
	w := doRequest(r, "GET", "/api/equipment/eq-race", nil)
	if w.Code != 200 {
		t.Fatalf("expected 200 after concurrent access, got %d", w.Code)
	}
}

func TestConcurrentSessionUpdates(t *testing.T) {
	r, _ := setupTestRouter()

	w := doRequest(r, "POST", "/api/sessions", nil)
	var sess model.Session
	parseJSON(w, &sess)

	// Concurrent viewport + highlights + occupancy updates
	done := make(chan bool, 3)
	go func() {
		for i := 0; i < 50; i++ {
			doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/viewport",
				model.Viewport{Room: "R001", Zoom: float64(i), Mode: "3d"})
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 50; i++ {
			doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/highlights",
				[]model.RoomHighlight{{RoomID: i % 3, Color: "#ff0000", Opacity: 0.5}})
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 50; i++ {
			occ := map[int]model.RoomOccupancy{
				0: {Persons: []model.Person{{ID: "p1", Name: "A"}}, Aliens: []model.Alien{}},
			}
			doRequest(r, "PUT", "/api/sessions/"+sess.ID+"/occupancy", occ)
		}
		done <- true
	}()
	for i := 0; i < 3; i++ {
		<-done
	}

	// Should still be consistent
	w = doRequest(r, "GET", "/api/sessions/"+sess.ID, nil)
	if w.Code != 200 {
		t.Fatalf("expected 200 after concurrent access, got %d", w.Code)
	}
}

func TestConcurrentEquipmentCreateDelete(t *testing.T) {
	r, _ := setupTestRouter()

	done := make(chan bool, 20)
	// 10 goroutines creating equipment
	for i := 0; i < 10; i++ {
		go func(n int) {
			id := "eq-conc-" + string(rune('a'+n))
			doRequest(r, "POST", "/api/equipment", model.Equipment{
				ID: id, Name: "Concurrent", Type: "test", Level: "level0", Room: "R001",
			})
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// 10 goroutines deleting
	for i := 0; i < 10; i++ {
		go func(n int) {
			id := "eq-conc-" + string(rune('a'+n))
			doRequest(r, "DELETE", "/api/equipment/"+id, nil)
			done <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	// Should be empty
	w := doRequest(r, "GET", "/api/equipment", nil)
	var list []*model.Equipment
	parseJSON(w, &list)
	if len(list) != 0 {
		t.Fatalf("expected 0 equipment after concurrent delete, got %d", len(list))
	}
}

func TestDeleteEquipmentCleansUpSensors(t *testing.T) {
	r, _ := setupTestRouter()

	doRequest(r, "POST", "/api/equipment", model.Equipment{
		ID: "eq-clean", Name: "Test", Type: "test", Level: "level0", Room: "R001",
	})
	doRequest(r, "POST", "/api/equipment/eq-clean/sensors",
		model.Sensor{ID: "s-clean", Name: "S", Type: "t", DataType: "text"})

	// Sensor should work
	w := doRequest(r, "PUT", "/api/sensors/s-clean/value",
		model.SensorValue{DataType: "text", Value: "1"})
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Delete equipment
	doRequest(r, "DELETE", "/api/equipment/eq-clean", nil)

	// Sensor should be gone
	w = doRequest(r, "PUT", "/api/sensors/s-clean/value",
		model.SensorValue{DataType: "text", Value: "2"})
	if w.Code != 404 {
		t.Fatalf("expected 404 after equipment delete, got %d", w.Code)
	}
}
