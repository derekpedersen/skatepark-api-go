package main

import (
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/api"
	"github.com/jeanphorn/log4go"
)

func main() {
	// configure logging
	log4go.LoadConfiguration("./log4go.json")

	// skateparks
	http.HandleFunc("/api/skatepark", api.SkateparkController)
	http.HandleFunc("/api/skatepark/states", api.SkateparkStatesController)

	// start server
	http.ListenAndServe(":8080", nil)
}
