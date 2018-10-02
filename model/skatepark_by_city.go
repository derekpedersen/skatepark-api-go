package model

// SkateparkByCity is a collection of skateparks in that city
type SkateparkByCity struct {
	City       string      `json:"city"`
	Skateparks []Skatepark `json:"skateparks"`
}
