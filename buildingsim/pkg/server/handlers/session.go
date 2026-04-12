package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type SessionHandlers struct {
	Store *store.MemoryStore
	Hub   *websocket.Hub
}

func (h *SessionHandlers) List(c *gin.Context) {
	sessions := h.Store.ListSessions()
	if sessions == nil {
		sessions = []*model.Session{}
	}
	c.JSON(http.StatusOK, sessions)
}

func (h *SessionHandlers) Create(c *gin.Context) {
	id := uuid.New().String()
	sess := h.Store.CreateSession(id)
	c.JSON(http.StatusCreated, sess)
}

func (h *SessionHandlers) Get(c *gin.Context) {
	id := c.Param("id")
	sess, ok := h.Store.GetSession(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	c.JSON(http.StatusOK, sess)
}

func (h *SessionHandlers) Delete(c *gin.Context) {
	id := c.Param("id")
	if !h.Store.DeleteSession(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *SessionHandlers) SetViewport(c *gin.Context) {
	id := c.Param("id")
	var vp model.Viewport
	if err := c.ShouldBindJSON(&vp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := h.Store.UpdateSessionViewport(id, vp); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	h.Hub.SendToSession(id, websocket.Message{Type: "viewport", Data: vp})
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SessionHandlers) SetHighlights(c *gin.Context) {
	id := c.Param("id")
	var highlights []model.RoomHighlight
	if err := c.ShouldBindJSON(&highlights); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := h.Store.UpdateSessionHighlights(id, highlights); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	h.Hub.SendToSession(id, websocket.Message{Type: "highlights", Data: highlights})
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SessionHandlers) SetOccupancy(c *gin.Context) {
	id := c.Param("id")
	var occupancy map[int]model.RoomOccupancy
	if err := c.ShouldBindJSON(&occupancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	version, ok := h.Store.UpdateSessionOccupancy(id, occupancy)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	h.Hub.SendToSession(id, websocket.Message{Type: "occupancy", Version: version})
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SessionHandlers) SetRoute(c *gin.Context) {
	id := c.Param("id")
	var route model.RouteResult
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !h.Store.UpdateSessionRoute(id, &route) {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	h.Hub.SendToSession(id, websocket.Message{Type: "route", Data: route})
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SessionHandlers) SetCoverage(c *gin.Context) {
	id := c.Param("id")
	var coverage []model.CoverageZone
	if err := c.ShouldBindJSON(&coverage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := h.Store.UpdateSessionCoverage(id, coverage); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	h.Hub.SendToSession(id, websocket.Message{Type: "coverage", Data: coverage})
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *SessionHandlers) HandleWebSocket(c *gin.Context) {
	id := c.Param("id")
	if _, ok := h.Store.GetSession(id); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	upgrader := websocket.Upgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	h.Hub.HandleConnection(conn, id)
}
