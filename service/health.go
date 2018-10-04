package service

import "github.com/derekpedersen/skatepark-api-go/model"

const (
	// HealthyMessage indicates the service is healthy
	HealthyMessage = "Johnny 5 Healthy!"

	// AliveMessage indicates the service is alive
	AliveMessage = "Johnny 5 Alive!"

	// ReadyMessage indicates the service is ready
	ReadyMessage = "Johnny 5 Ready!"
)

// HealthService interface
type HealthService interface {
	GetHealthyMessage() *model.HealthMessage
	GetAliveMessage() *model.HealthMessage
	GetReadyMessage() *model.HealthMessage
}

// HealthServiceImpl struct
type HealthServiceImpl struct {
}

// NewHealthService creates a new HealthService
func NewHealthService() HealthService {
	return &HealthServiceImpl{}
}

// GetHealthyMessage gets the alive message
func (a *HealthServiceImpl) GetHealthyMessage() *model.HealthMessage {
	return model.NewHealthMessage(HealthyMessage)
}

// GetAliveMessage gets the alive message
func (a *HealthServiceImpl) GetAliveMessage() *model.HealthMessage {
	return model.NewHealthMessage(AliveMessage)
}

// GetReadyMessage gets the ready message
func (a *HealthServiceImpl) GetReadyMessage() *model.HealthMessage {
	return model.NewHealthMessage(ReadyMessage)
}
