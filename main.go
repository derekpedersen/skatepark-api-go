//go:generate swagger generate spec -m -o .docs/swagger/swagger.json

// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/jsons
//
// swagger:meta
package main

import (
	"net/http"
	"os"
	"sync"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/appcfg"
	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel) // TODO: make this a flag
	log.Infof("log level: %v", log.GetLevel())

	// setup imgur services
	apiKey := os.Getenv("IMGUR_API_KEY")
	imgSvc := imgurService.NewAlbumService(apiKey)

	// setup skatepark services
	skRepo := repository.NewSkateparkRepository()
	skSvc := service.NewSkateparksService(imgSvc, skRepo)
	skCtrl := controller.NewSkateparksAPIController(skSvc)

	// setup health services
	hsvc := service.NewHealthService()
	hctrl := controller.NewHealthAPIController(hsvc, skSvc)

	// setup api routers
	baseRouter, err := appcfg.NewBaseRouter()
	if err != nil {
		log.Fatalf("failed to create baseRouter: %v", err)
	}
	appcfg.AddHealthRoutes(baseRouter, hctrl)
	skateparkRouter, err := appcfg.NewSkateparkAPIRouter(baseRouter, skCtrl)
	if err != nil {
		log.Fatalf("failed to create skateparkRouter: %v", err)
	}
	swaggerRouter := appcfg.NewSwaggerRouter("./.docs/swagger/") // TODO: this should be an env var

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
		// Debug: true,
	})

	var wg sync.WaitGroup

	// start the api server
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":8080", c.Handler(skateparkRouter)))
	}()

	// start the swagger server
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":3000", c.Handler(swaggerRouter)))
	}()

	wg.Wait()
}
