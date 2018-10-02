package model

// SkateparkByState is a collection of skateparks grouped by state then city
type SkateparkByState struct {
	State  string            `json:"state"`
	Cities []SkateparkByCity `json:"cities"`
}
