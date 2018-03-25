package controller

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/jeanphorn/log4go"
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

// GetSkateparks gets the full collection of skateparks
func (api *SkateparksAPIControllerImpl) GetSkateparks(w http.ResponseWriter, r *http.Request) {
	skateparks, err := api.svc.GetSkateparks()
	if err != nil {
		log4go.Error("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log4go.Info("Number of Skateparks: %v", len(skateparks))

	sortedBy := r.URL.Query().Get("sortedBy")
	var js []byte
	switch sortedBy {
	case "state":
		js, err = json.Marshal(skateparks.GetSkateparksByState())
	default:
		js, err = json.Marshal(skateparks)
	}

	if err != nil {
		log4go.Error("Error in marshalling skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetSkateparksByState gets the full collection of skateparks
func (api *SkateparksAPIControllerImpl) GetSkateparksByState(w http.ResponseWriter, r *http.Request) {
	skateparks, err := api.svc.GetSkateparks()
	if err != nil {
		log4go.Error("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log4go.Info("Number of Skateparks: %v", len(skateparks))

	js, err := json.Marshal(skateparks.GetSkateparksByState())

	if err != nil {
		log4go.Error("Error in marshalling skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
