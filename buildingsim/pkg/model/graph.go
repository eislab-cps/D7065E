package model

type NavGraph struct {
	Nodes []NavNode `json:"nodes"`
	Edges []NavEdge `json:"edges"`
}

type NavNode struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Type  string  `json:"type,omitempty"`  // "corridor", "entry", "room" (walkable graph)
	Level string  `json:"level,omitempty"` // floor level (multi-floor graph)
}

type NavEdge struct {
	From   int     `json:"from"`
	To     int     `json:"to"`
	Weight float64 `json:"weight"`
	X1     float64 `json:"x1,omitempty"`
	Y1     float64 `json:"y1,omitempty"`
	X2     float64 `json:"x2,omitempty"`
	Y2     float64 `json:"y2,omitempty"`
}

type RouteResult struct {
	Path     []RouteNode `json:"path"`
	Distance float64     `json:"distance"`
}

type RouteNode struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Level  string `json:"level,omitempty"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}
