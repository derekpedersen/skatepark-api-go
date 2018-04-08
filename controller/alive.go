package controller

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/service"
)

// Controller
func AliveController(w http.ResponseWriter, r *http.Request) {
	alive := service.GetAliveMessage()

	js, err := json.Marshal(*alive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
