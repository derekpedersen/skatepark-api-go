package model

// SkateparkByState
type SkateparkByState struct {
	State  string            `json:"state"`
	Cities []SkateparkByCity `json:"cities"`
}
