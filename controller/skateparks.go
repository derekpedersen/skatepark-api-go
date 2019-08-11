package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/derekpedersen/skatepark-api-go/service"
	log "github.com/sirupsen/logrus"
)

// SkateparksAPIController interface
type SkateparksAPIController interface {
	GetSkateparks(w http.ResponseWriter, r *http.Request)
	GetSkateparksByState(w http.ResponseWriter, r *http.Request)
}

// SkateparksAPIControllerImpl implementation
type SkateparksAPIControllerImpl struct {
	svc service.SkateparksService
}

// NewSkateparksAPIController creates a new skateparks controller
func NewSkateparksAPIController(svc service.SkateparksService) *SkateparksAPIControllerImpl {
	return &SkateparksAPIControllerImpl{
		svc: svc,
	}
}

// GetSkateparksByCity gets the full collection of skateparks
func (api *SkateparksAPIControllerImpl) GetSkateparksByCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	skateparks, err := api.svc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Number of Skateparks: %v", len(skateparks))

	m := skateparks.StateSkateparkMap()
	state := mux.Vars(r)["state"]
	if val, ok := m[state]; ok {
		js, err := json.Marshal(val)
		if err != nil {
			log.Errorf("Error in marshalling skateparks:\n %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	} else {
		http.NotFound(w, r)
	}
	return
}

// GetSkateparksByState gets the full collection of skateparks
func (api *SkateparksAPIControllerImpl) GetSkateparksByState(w http.ResponseWriter, r *http.Request) {
	skateparks, err := api.svc.GetSkateparks()
	if err != nil {
		log.Errorf("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Infof("Number of Skateparks: %v", len(skateparks))

	js, err := json.Marshal(skateparks.StateSkateparkMap())

	if err != nil {
		log.Errorf("Error in marshalling skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
