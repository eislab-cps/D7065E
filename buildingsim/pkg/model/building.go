package model

type Building struct {
	Name   string  `json:"name"`
	Levels []Level `json:"levels"`
}

type Level struct {
	ID    string `json:"id"`    // e.g. "level0"
	Label string `json:"label"` // e.g. "Floor 0"
}

type FloorData struct {
	Page       Page        `json:"page"`
	Rooms      []Room      `json:"rooms"`
	Walls      [][][2]float64 `json:"walls"`
	RedLines   [][][2]float64 `json:"red_lines"`
	GreenLines [][][2]float64 `json:"green_lines"`
	Labels     []Label     `json:"labels"`
	Graph         *NavGraph `json:"graph,omitempty"`
	WalkableGraph *NavGraph `json:"walkable_graph,omitempty"`
}

type Page struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Room struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Area    float64      `json:"area"`
	Center  [2]float64   `json:"center"`
	Polygon [][2]float64 `json:"polygon"`
	Type    string       `json:"type,omitempty"`
}

type Label struct {
	Text string  `json:"text"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

type CrossFloorEdge struct {
	FromLevel string `json:"from_level"`
	FromName  string `json:"from_name"`
	FromID    int    `json:"from_id"`
	ToLevel   string `json:"to_level"`
	ToName    string `json:"to_name"`
	ToID      int    `json:"to_id"`
	Connector string `json:"connector"`
	Type      string `json:"type"` // "stair" or "elevator"
}
