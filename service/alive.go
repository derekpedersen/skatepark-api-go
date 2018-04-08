package service

import "github.com/derekpedersen/skatepark-api-go/model"

// GetAliveMessage
func GetAliveMessage() *model.Alive {
	return model.NewAlive("Johnny 5 Alive!")
}
