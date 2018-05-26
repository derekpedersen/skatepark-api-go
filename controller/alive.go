package controller

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
)

type AliveAPIController interface {
	GetAliveMessage(w http.ResponseWriter, r *http.Request)
}

type AliveAPIControllerImpl struct {
	svc service.AliveService
}

func NewAliveAPIController(svc service.AliveService) *AliveAPIControllerImpl {
	return &AliveAPIControllerImpl{
		svc: svc,
	}
}

// GetAliveMessage controller
func (api *AliveAPIControllerImpl) GetAliveMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetAliveMessage()

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
