package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/server/websocket"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type EquipmentHandlers struct {
	Store *store.MemoryStore
	Hub   *websocket.Hub
}

func (h *EquipmentHandlers) Create(c *gin.Context) {
	var eq model.Equipment
	if err := c.ShouldBindJSON(&eq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if eq.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	if _, exists := h.Store.GetEquipment(eq.ID); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "equipment already exists"})
		return
	}
	h.Store.CreateEquipment(&eq)
	c.JSON(http.StatusCreated, eq)
}

func (h *EquipmentHandlers) List(c *gin.Context) {
	level := c.Query("level")
	room := c.Query("room")
	typ := c.Query("type")
	category := c.Query("category")
	result := h.Store.ListEquipment(level, room, typ, category)
	if result == nil {
		result = []*model.Equipment{}
	}
	c.JSON(http.StatusOK, result)
}

func (h *EquipmentHandlers) Get(c *gin.Context) {
	id := c.Param("id")
	eq, ok := h.Store.GetEquipment(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "equipment not found"})
		return
	}
	c.JSON(http.StatusOK, eq)
}

func (h *EquipmentHandlers) Update(c *gin.Context) {
	id := c.Param("id")
	var eq model.Equipment
	if err := c.ShouldBindJSON(&eq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	eq.ID = id
	if !h.Store.UpdateEquipment(&eq) {
		c.JSON(http.StatusNotFound, gin.H{"error": "equipment not found"})
		return
	}
	c.JSON(http.StatusOK, eq)
}

func (h *EquipmentHandlers) Delete(c *gin.Context) {
	id := c.Param("id")
	if !h.Store.DeleteEquipment(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "equipment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *EquipmentHandlers) BulkCreate(c *gin.Context) {
	var items []model.Equipment
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created := 0
	for i := range items {
		eq := &items[i]
		if eq.ID == "" {
			continue
		}
		if _, exists := h.Store.GetEquipment(eq.ID); exists {
			continue
		}
		h.Store.CreateEquipment(eq)
		created++
	}
	version := h.Store.BumpEquipmentVersion()
	h.Hub.BroadcastToAll(websocket.Message{Type: "equipment", Version: version})
	c.JSON(http.StatusCreated, gin.H{"created": created, "total": len(items), "version": version})
}

func (h *EquipmentHandlers) Notify(c *gin.Context) {
	version := h.Store.BumpEquipmentVersion()
	h.Hub.BroadcastToAll(websocket.Message{
		Type:    "equipment",
		Version: version,
	})
	c.JSON(http.StatusOK, gin.H{"version": version})
}
