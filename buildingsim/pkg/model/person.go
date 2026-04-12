package model

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon,omitempty"` // "man", "woman", or empty (defaults to "man")
}

type Alien struct {
	ID string `json:"id"`
}
