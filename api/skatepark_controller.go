package api

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/jeanphorn/log4go"
)

// SkateparkController gets the full collection of skateparks
func SkateparkController(w http.ResponseWriter, r *http.Request) {
	skateparks, err := service.GetSkateparks()
	if err != nil {
		log4go.Error("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log4go.Info("Number of Skateparks: %v", len(skateparks))

	js, err := json.Marshal(skateparks)
	if err != nil {
		log4go.Error("Error in marshalling skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
