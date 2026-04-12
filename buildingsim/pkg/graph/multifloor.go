package graph

import (
	"github.com/eislab-cps/buildingsim/pkg/model"
)

// BuildMultiFloorGraph merges per-floor walkable graphs into a single graph
// with cross-floor edges connecting stairs/elevators.
func BuildMultiFloorGraph(
	floors map[string]*model.FloorData,
	levelOrder []string,
	crossFloorEdges []model.CrossFloorEdge,
) *model.NavGraph {
	merged := &model.NavGraph{}
	nextID := 0

	// Map: level + original_node_id -> new_node_id
	idMap := make(map[string]map[int]int) // level -> old_id -> new_id
	// Map: level + room_name -> new_node_id (for room-type nodes)
	nameMap := make(map[string]map[string]int) // level -> name -> new_id

	for _, level := range levelOrder {
		data, ok := floors[level]
		if !ok || data.WalkableGraph == nil {
			continue
		}
		wg := data.WalkableGraph
		idMap[level] = make(map[int]int)
		nameMap[level] = make(map[string]int)

		// Add nodes with new IDs and level annotation
		for _, node := range wg.Nodes {
			newID := nextID
			nextID++
			idMap[level][node.ID] = newID
			merged.Nodes = append(merged.Nodes, model.NavNode{
				ID:   newID,
				Name: node.Name,
				X:    node.X,
				Y:    node.Y,
				Type: node.Type,
				Level: level,
			})
			// Track room and entry nodes by name for cross-floor linking
			if node.Type == "room" || node.Type == "entry" {
				// Prefer entry nodes (closer to corridor) for cross-floor connections
				if _, exists := nameMap[level][node.Name]; !exists || node.Type == "entry" {
					nameMap[level][node.Name] = newID
				}
			}
		}

		// Add edges with remapped IDs
		for _, edge := range wg.Edges {
			newFrom, ok1 := idMap[level][edge.From]
			newTo, ok2 := idMap[level][edge.To]
			if !ok1 || !ok2 {
				continue
			}
			merged.Edges = append(merged.Edges, model.NavEdge{
				From:   newFrom,
				To:     newTo,
				Weight: edge.Weight,
				X1:     edge.X1,
				Y1:     edge.Y1,
				X2:     edge.X2,
				Y2:     edge.Y2,
			})
		}
	}

	// Add cross-floor edges
	for _, cfe := range crossFloorEdges {
		fromLevel := extractLevel(cfe.FromLevel)
		toLevel := extractLevel(cfe.ToLevel)

		fromNames, ok1 := nameMap[fromLevel]
		toNames, ok2 := nameMap[toLevel]
		if !ok1 || !ok2 {
			continue
		}

		fromID, ok1 := fromNames[cfe.FromName]
		toID, ok2 := toNames[cfe.ToName]
		if !ok1 || !ok2 {
			continue
		}

		// Cross-floor weight: fixed cost for stairs/elevator
		weight := 20.0
		if cfe.Type == "elevator" {
			weight = 10.0
		}

		merged.Edges = append(merged.Edges, model.NavEdge{
			From:   fromID,
			To:     toID,
			Weight: weight,
		})
	}

	return merged
}

func extractLevel(s string) string {
	// "abuilding/level0" -> "level0"
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			return s[i+1:]
		}
	}
	return s
}
