package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type BuildingHandlers struct {
	Store *store.MemoryStore
}

func (h *BuildingHandlers) GetBuilding(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetBuilding())
}

func (h *BuildingHandlers) GetFloor(c *gin.Context) {
	level := c.Param("level")
	data, ok := h.Store.GetFloorData(level)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "level not found"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *BuildingHandlers) GetCrossFloorEdges(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.GetCrossFloorEdges())
}
