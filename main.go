package main

import (
	"net/http"
	"sync"

	"github.com/derekpedersen/skatepark-api-go/appcfg"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel) // TODO: make this a flag
	log.Infof("log level: %v", log.GetLevel())

	var wg sync.WaitGroup

	skateparkRouter, _ := appcfg.NewSkateparkAPIRouter()

	// Use default options
	handler := cors.Default().Handler(skateparkRouter)

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8080", handler))
	}()

	wg.Wait()
}
