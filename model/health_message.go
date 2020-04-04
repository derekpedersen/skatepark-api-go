package model

// HealthMessage struct represents the alive messaging
// swagger:model healthMessage
type HealthMessage struct {
	Message string `json:"message"`
}

// NewHealthMessage creates a new HealthMessage struct
func NewHealthMessage(str string) *HealthMessage {
	return &HealthMessage{
		Message: str,
	}
}
