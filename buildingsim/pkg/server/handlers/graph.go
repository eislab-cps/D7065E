package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/eislab-cps/buildingsim/pkg/graph"
	"github.com/eislab-cps/buildingsim/pkg/model"
	"github.com/eislab-cps/buildingsim/pkg/store"
)

type GraphHandlers struct {
	Store *store.MemoryStore
}

func (h *GraphHandlers) getNavGraph(c *gin.Context, data *model.FloorData) *model.NavGraph {
	graphType := c.DefaultQuery("type", "adjacency")
	if graphType == "walkable" {
		return data.WalkableGraph
	}
	return data.Graph
}

func (h *GraphHandlers) GetGraph(c *gin.Context) {
	level := c.DefaultQuery("level", "level0")
	data, ok := h.Store.GetFloorData(level)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "level not found"})
		return
	}
	g := h.getNavGraph(c, data)
	if g == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "graph not found"})
		return
	}
	c.JSON(http.StatusOK, g)
}

func (h *GraphHandlers) GetRoute(c *gin.Context) {
	level := c.Query("level")
	graphType := c.DefaultQuery("type", "adjacency")

	var result *model.RouteResult

	fromName := c.Query("from_name")
	toName := c.Query("to_name")

	if fromName != "" && toName != "" && graphType == "walkable" && level == "" {
		// Multi-floor walkable routing
		g := h.Store.GetMultiFloorGraph()
		if g == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "multi-floor graph not available"})
			return
		}
		result = graph.ShortestPathByName(g, fromName, toName)
	} else {
		// Single-floor routing
		if level == "" {
			level = "level0"
		}
		data, ok := h.Store.GetFloorData(level)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "level not found"})
			return
		}
		g := h.getNavGraph(c, data)
		if g == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "graph not found"})
			return
		}

		if fromName != "" && toName != "" {
			result = graph.ShortestPathByName(g, fromName, toName)
		} else {
			fromStr := c.Query("from")
			toStr := c.Query("to")
			from, err := strconv.Atoi(fromStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'from' room id"})
				return
			}
			to, err := strconv.Atoi(toStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'to' room id"})
				return
			}
			result = graph.ShortestPath(g, from, to)
		}
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no route found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
