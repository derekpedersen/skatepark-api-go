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
//	Schemes: http, https
//	Host: localhost
//	BasePath: /v2
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/jsons
//
// swagger:meta
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/derekpedersen/imgur-go/album"
	"github.com/derekpedersen/imgur-go/authorization"
	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func parseBoolEnvVar(
	envVar string,
) (
	result bool,
) {

	result, err := strconv.ParseBool(envVar)
	if err != nil {
		return false
	}

	return result
}

var LOAD_IMGUR_ALBUMS = flag.Bool("load-imgur-albums", parseBoolEnvVar(os.Getenv("LOAD_IMGUR_ALBUMS")), "Load Imgur Albums controls if albums are refreshed")

func main() {

	log.Debug("Intializing application...")
	log.SetLevel(log.DebugLevel) // TODO: make this a flag
	log.Infof("log level: %v", log.GetLevel())

	log.Debug("Intializing imgur service...")
	imgurAuth, err := authorization.NewAuthorization()
	if err != nil {
		log.Fatal(err)
	}

	if skatepark_api.IMGUR_SVC = album.NewAlbumService(*imgurAuth, "https://api.imgur.com/3/album/"); skatepark_api.IMGUR_SVC == nil {
		log.Fatal(fmt.Errorf("failed to load imgur service"))
	}

	log.Debug("Loading skateparks db...")
	if _, err := skatepark_api.GetSkateparks(true, *LOAD_IMGUR_ALBUMS); err != nil {
		log.Fatalf("failed to parse skateparks: %v", err)
	}

	log.Debug("Intializing health service...")
	hsvc := skatepark_api.NewHealthService()
	hctrl := skatepark_api.NewHealthAPIController(hsvc)

	log.Debug("Intializing api router...")
	baseRouter, err := skatepark_api.NewBaseRouter()
	if err != nil {
		log.Fatalf("failed to create api base router: %v", err)
	}

	log.Debug("Adding health routes to api router...")
	skatepark_api.AddHealthRoutes(baseRouter, hctrl)

	log.Debug("Adding skatepark routes to api router...")
	skateparkRouter, err := skatepark_api.NewSkateparkAPIRouter(baseRouter)
	if err != nil {
		log.Fatalf("failed to add skatepark api routes: %v", err)
	}

	log.Debug("Initialzing swagger router...")
	swaggerRouter := skatepark_api.NewSwaggerRouter("./.docs/swagger/") // TODO: this should be an env var

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
