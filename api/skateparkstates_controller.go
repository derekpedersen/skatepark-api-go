package api

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/jeanphorn/log4go"
)

// StatesController gets a collection of skateparks grouped by state
func SkateparkStatesController(w http.ResponseWriter, r *http.Request) {
	skateparks, err := service.GetSkateparks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log4go.Info("Number of States: %v", len(skateparks))

	sk := service.Skateparks(skateparks)
	states := sk.GetSkateparksByState()

	js, err := json.Marshal(states)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
