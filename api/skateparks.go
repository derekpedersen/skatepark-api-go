package api

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/jeanphorn/log4go"
)

// SkateparkController gets the full collection of skateparks
func SkateparkController(w http.ResponseWriter, r *http.Request) {
	sk, err := service.GetSkateparks()
	if err != nil {
		log4go.Error("Error in getting skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	skateparks := service.Skateparks(sk)
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
