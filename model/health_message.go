package model

// HealthMessage struct represents the alive messaging
// swagger:model
type HealthMessage struct {
	Message string `json:"message"`
}

// NewHealthMessage creates a new HealthMessage struct
// swagger:model
func NewHealthMessage(str string) *HealthMessage {
	return &HealthMessage{
		Message: str,
	}
}
