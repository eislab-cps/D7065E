package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type CoverageHandlers struct {
	Store *store.MemoryStore
	Hub   *websocket.Hub
}

func (h *CoverageHandlers) Get(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetCoverage())
}

func (h *CoverageHandlers) Set(c *gin.Context) {
	var zones []model.CoverageZone
	if err := c.ShouldBindJSON(&zones); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	version := h.Store.SetCoverage(zones)
	h.Hub.BroadcastToAll(websocket.Message{Type: "coverage", Data: zones, Version: version})
	c.JSON(http.StatusOK, gin.H{"status": "updated", "version": version})
}
