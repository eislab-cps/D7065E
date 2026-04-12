package model

import "time"

type Session struct {
	ID           string                  `json:"id"`
	Viewport     Viewport                `json:"viewport"`
	Highlights   []RoomHighlight         `json:"highlights"`
	Occupancy    map[int]RoomOccupancy   `json:"occupancy"`    // deprecated, use global /api/occupancy
	Route        *RouteResult            `json:"route,omitempty"`
	Coverage     []CoverageZone          `json:"coverage"`
	LastWSActive time.Time               `json:"last_ws_active"`
	Version      int64                   `json:"version"`
}

type CoverageZone struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Room     string     `json:"room,omitempty"`  // room name to center on (resolved by browser)
	Center   [2]float64 `json:"center"`          // [x, y] in PDF coordinates (used if room is empty)
	Radius   float64    `json:"radius"`          // radius in PDF units
	Color    string     `json:"color"`           // hex color
	Opacity  float64    `json:"opacity"`         // 0.0 - 1.0
	Level    string     `json:"level"`           // "level0", "level1", "level2", or "" for all floors
	Height   float64    `json:"height"`          // vertical extent (3D), 0 = flat disc
}

type Viewport struct {
	Room    string  `json:"room,omitempty"`    // room name to center on (e.g. "A2306")
	Zoom    float64 `json:"zoom"`
	Mode    string  `json:"mode"`              // "2d" or "3d"
	Floor   string  `json:"floor,omitempty"`   // "level0", "level1", "level2"
}

type RoomHighlight struct {
	RoomID  int     `json:"room_id"`
	Color   string  `json:"color"`
	Opacity float64 `json:"opacity"`
}

type RoomOccupancy struct {
	Persons []Person `json:"persons"`
	Aliens  []Alien  `json:"aliens"`
}
