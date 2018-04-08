package main

import (
	"net/http"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/jeanphorn/log4go"
)

func main() {
	// configure logging
	log4go.LoadConfiguration("./log4go.json")

	// alive
	http.HandleFunc("/alive", controller.AliveController)

	// skateparks
	http.HandleFunc("/api/skatepark", controller.SkateparkController)
	http.HandleFunc("/api/skatepark/states", controller.SkateparkStatesController)

	// start server
	http.ListenAndServe(":8080", nil)
}
