package model

type SkateparkByCity struct {
	City       string      `json:"city"`
	Skateparks []Skatepark `json:"skateparks"`
}
