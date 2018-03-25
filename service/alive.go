package service

import "github.com/derekpedersen/skatepark-api-go/model"

type AliveService interface {
	GetAliveMessage() *model.Alive
}

type AliveServiceImpl struct {
}

func NewAliveService() *AliveServiceImpl {
	return &AliveServiceImpl{}
}

// GetAliveMessage gets the alive
func (a *AliveServiceImpl) GetAliveMessage() *model.Alive {
	return model.NewAlive("Johnny 5 Alive!")
}
