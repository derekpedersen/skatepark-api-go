package api

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
)

// StatesController gets a collection of skateparks grouped by state
func SkateparkStatesController(w http.ResponseWriter, r *http.Request) {
	skateparks, err := service.GetSkateparksByState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(skateparks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
