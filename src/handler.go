package skatepark_api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// GetSkateparksByState gets the full collection of skateparks
func GetSkateparksByState(w http.ResponseWriter, r *http.Request) {

	m := SKATEPARKS.StateSkateparkMap()

	state := mux.Vars(r)["state"]
	if len(state) > 0 {
		if val, ok := m[state]; ok {
			js, err := json.Marshal(val)
			if err != nil {
				logrus.Errorf("Error in marshalling skateparks:\n %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			return
		}
		http.NotFound(w, r)
		return
	}

	js, err := json.Marshal(SKATEPARKS.StateSkateparkMap())

	if err != nil {
		log.Errorf("Error in marshalling skateparks:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetSkateparksByCity gets the full collection of skateparks
func GetSkateparksByCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := SKATEPARKS.CitySkateparkMap()
	city := mux.Vars(r)["city"]
	if val, ok := m[city]; ok {
		js, err := json.Marshal(val)
		if err != nil {
			logrus.Errorf("Error in marshalling skateparks:\n %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	} else {
		http.NotFound(w, r)
	}
	return
}

// GetSkateparksByName gets the full collection of skateparks
func GetSkateparksByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := mux.Vars(r)["skatepark"]

	s := SKATEPARKS.GetSkateparkByName(name)
	js, err := json.Marshal(s)
	if err != nil {
		logrus.Errorf("Error in marshalling skatepark:\n %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
	return
}
