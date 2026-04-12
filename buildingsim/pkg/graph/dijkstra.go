package graph

import (
	"container/heap"
	"math"

	"github.com/eislab-cps/buildingsim/pkg/model"
)

type adjacency struct {
	to     int
	weight float64
}

type item struct {
	node int
	dist float64
	idx  int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].idx = i; pq[j].idx = j }
func (pq *priorityQueue) Push(x interface{}) { it := x.(*item); it.idx = len(*pq); *pq = append(*pq, it) }
func (pq *priorityQueue) Pop() interface{}   { old := *pq; n := len(old); it := old[n-1]; *pq = old[:n-1]; return it }

// ShortestPath computes Dijkstra's shortest path between two room IDs on a single floor.
func ShortestPath(graph *model.NavGraph, fromID, toID int) *model.RouteResult {
	if graph == nil {
		return nil
	}

	nodeIndex := make(map[int]int) // room id -> index in nodes
	for i, n := range graph.Nodes {
		nodeIndex[n.ID] = i
	}

	if _, ok := nodeIndex[fromID]; !ok {
		return nil
	}
	if _, ok := nodeIndex[toID]; !ok {
		return nil
	}

	// Build adjacency list
	n := len(graph.Nodes)
	adj := make([][]adjacency, n)
	for i := range adj {
		adj[i] = []adjacency{}
	}
	for _, e := range graph.Edges {
		fi, ok1 := nodeIndex[e.From]
		ti, ok2 := nodeIndex[e.To]
		if !ok1 || !ok2 {
			continue
		}
		adj[fi] = append(adj[fi], adjacency{to: ti, weight: e.Weight})
		adj[ti] = append(adj[ti], adjacency{to: fi, weight: e.Weight})
	}

	// Dijkstra
	dist := make([]float64, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = math.Inf(1)
		prev[i] = -1
	}

	startIdx := nodeIndex[fromID]
	endIdx := nodeIndex[toID]
	dist[startIdx] = 0

	pq := &priorityQueue{&item{node: startIdx, dist: 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*item)
		if cur.dist > dist[cur.node] {
			continue
		}
		if cur.node == endIdx {
			break
		}
		for _, edge := range adj[cur.node] {
			newDist := dist[cur.node] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				prev[edge.to] = cur.node
				heap.Push(pq, &item{node: edge.to, dist: newDist})
			}
		}
	}

	if math.IsInf(dist[endIdx], 1) {
		return nil
	}

	// Reconstruct path
	var path []model.RouteNode
	for idx := endIdx; idx != -1; idx = prev[idx] {
		node := graph.Nodes[idx]
		path = append([]model.RouteNode{{
			RoomID: node.ID,
			Name:   node.Name,
			Level:  node.Level,
			X:      node.X,
			Y:      node.Y,
		}}, path...)
	}

	return &model.RouteResult{
		Path:     path,
		Distance: dist[endIdx],
	}
}

// ShortestPathByName finds nodes by name (room label) and computes shortest path.
// For walkable graphs, multiple nodes may share the same name (room node + entry node).
// We pick the "room" type node if available.
func ShortestPathByName(graph *model.NavGraph, fromName, toName string) *model.RouteResult {
	if graph == nil {
		return nil
	}

	findNode := func(name string) int {
		bestID := -1
		for _, n := range graph.Nodes {
			if n.Name == name {
				if n.Type == "room" || bestID == -1 {
					bestID = n.ID
				}
			}
		}
		return bestID
	}

	fromID := findNode(fromName)
	toID := findNode(toName)
	if fromID == -1 || toID == -1 {
		return nil
	}
	return ShortestPath(graph, fromID, toID)
}
