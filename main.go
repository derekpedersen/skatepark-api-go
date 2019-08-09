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

	skateparkRouter, err := appcfg.NewSkateparkAPIRouter()
	if err != nil {
		log.Fatalf("failed to create skateparkRouter: %v", err)
	}

	// setup cors option
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost",
			"http://localhost:4200",
			"http://localhost:8080",
			"http://skatepark-api",
			"http://skatepark-api:8080",
			"http://skatepark-api-go-service",
			"http://skatepark-api-go-service:8080",
			"http://celebrityskateboards.com",
			"https://celebrityskateboards.com",
		},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		//Debug: true,
	})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8080", c.Handler(skateparkRouter)))
	}()

	wg.Wait()
}
