package model

// SkateparkByState
type SkateparkByState struct {
	State  string `json:"state"`
	Cities []struct {
		City       string      `json:"city"`
		Skateparks []Skatepark `json:"skateparks"`
	} `json:"cities"`
}
