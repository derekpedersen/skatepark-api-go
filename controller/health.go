package controller

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
)

// HealthAPIController interface
type HealthAPIController interface {
	GetAliveMessage(w http.ResponseWriter, r *http.Request)
	GetReadyMessage(w http.ResponseWriter, r *http.Request)
	GetHealthyMessage(w http.ResponseWriter, r *http.Request)
}

// HealthAPIControllerImpl struct
type HealthAPIControllerImpl struct {
	svc service.HealthService
}

// NewHealthAPIController creates a new HealthAPIController
func NewHealthAPIController(svc service.HealthService) HealthAPIController {
	return &HealthAPIControllerImpl{
		svc: svc,
	}
}

// GetAliveMessage controller
func (api *HealthAPIControllerImpl) GetAliveMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetAliveMessage()

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetReadyMessage controller
func (api *HealthAPIControllerImpl) GetReadyMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetReadyMessage()

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetHealthyMessage controller
func (api *HealthAPIControllerImpl) GetHealthyMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetHealthyMessage()

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
