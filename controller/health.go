package controller

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
	log "github.com/sirupsen/logrus"
)

// HealthAPIController interface
type HealthAPIController interface {
	GetAliveMessage(w http.ResponseWriter, r *http.Request)
	GetReadyMessage(w http.ResponseWriter, r *http.Request)
	GetHealthyMessage(w http.ResponseWriter, r *http.Request)
}

// HealthAPIControllerImpl struct
type HealthAPIControllerImpl struct {
	svc           service.HealthService
	skateparksSvc service.SkateparksService
}

// NewHealthAPIController creates a new HealthAPIController
func NewHealthAPIController(svc service.HealthService, skateparksSvc service.SkateparksService) HealthAPIController {
	return &HealthAPIControllerImpl{
		svc:           svc,
		skateparksSvc: skateparksSvc,
	}
}

// GetAliveMessage controller
func (api *HealthAPIControllerImpl) GetAliveMessage(w http.ResponseWriter, r *http.Request) {
	alive := api.svc.GetAliveMessage()

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	_, err := api.skateparksSvc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
