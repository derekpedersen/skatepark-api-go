package model

// Alive struct represents the alive messaging
type Alive struct {
	Message string `json:"message"`
}

// NewAlive creates a new Alive struct
func NewAlive(str string) *Alive {
	return &Alive{
		Message: str,
	}
}
