package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type OccupancyHandlers struct {
	Store *store.MemoryStore
	Hub   *websocket.Hub
}

func (h *OccupancyHandlers) Get(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetOccupancy())
}

func (h *OccupancyHandlers) Set(c *gin.Context) {
	var occ map[int]model.RoomOccupancy
	if err := c.ShouldBindJSON(&occ); err != nil {
		// Empty body or {} is valid — treat as clear
		occ = make(map[int]model.RoomOccupancy)
	}
	if occ == nil {
		occ = make(map[int]model.RoomOccupancy)
	}
	version := h.Store.SetOccupancy(occ)
	h.Hub.BroadcastToAll(websocket.Message{Type: "occupancy", Version: version})
	c.JSON(http.StatusOK, gin.H{"status": "updated", "version": version})
}
