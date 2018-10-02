package service

import "github.com/derekpedersen/skatepark-api-go/model"

// HealthService interface
type HealthService interface {
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

// GetAliveMessage gets the alive message
func (a *HealthServiceImpl) GetAliveMessage() *model.HealthMessage {
	return model.NewHealthMessage("Johnny 5 Alive!")
}

// GetReadyMessage gets the ready message
func (a *HealthServiceImpl) GetReadyMessage() *model.HealthMessage {
	return model.NewHealthMessage("Johnny 5 Ready!")
}
