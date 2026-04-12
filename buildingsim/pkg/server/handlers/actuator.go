package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type ActuatorHandlers struct {
	Store *store.MemoryStore
}

func (h *ActuatorHandlers) AddActuator(c *gin.Context) {
	equipmentID := c.Param("id")
	var act model.Actuator
	if err := c.ShouldBindJSON(&act); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if act.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	ok, errMsg := h.Store.AddActuator(equipmentID, &act)
	if !ok {
		if errMsg == "actuator already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": errMsg})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
		}
		return
	}
	c.JSON(http.StatusCreated, act)
}

func (h *ActuatorHandlers) ListActuators(c *gin.Context) {
	equipmentID := c.Param("id")
	actuators, ok := h.Store.GetActuatorsForEquipment(equipmentID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "equipment not found"})
		return
	}
	c.JSON(http.StatusOK, actuators)
}

func (h *ActuatorHandlers) DeleteActuator(c *gin.Context) {
	actuatorID := c.Param("id")
	if !h.Store.DeleteActuator(actuatorID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "actuator not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *ActuatorHandlers) SetState(c *gin.Context) {
	actuatorID := c.Param("id")
	var state model.ActuatorState
	if err := c.ShouldBindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !h.Store.SetActuatorState(actuatorID, state) {
		c.JSON(http.StatusNotFound, gin.H{"error": "actuator not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
