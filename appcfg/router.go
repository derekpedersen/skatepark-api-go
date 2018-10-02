package appcfg

import (
	"os"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/gorilla/mux"
)

// NewSkateparkAPIRouter creates a new mux router for the skatepark api
func NewSkateparkAPIRouter() (*mux.Router, error) {
	// setup imgur service
	apiKey := os.Getenv("IMGUR_API_KEY")
	imgSvc := imgurService.NewAlbumService(apiKey)

	// set "api" path prefix
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()

	skRepo := repository.NewSkateparkRepository()
	skSvc := service.NewSkateparksService(imgSvc, skRepo)
	skCtrl := controller.NewSkateparksAPIController(skSvc)
	router.HandleFunc("/skateparks", skCtrl.GetSkateparks)
	router.HandleFunc("/skatepark/states", skCtrl.GetSkateparksByState) // TODO: this is the legacy one, but need to update SPA first before removing

	return router, nil
}

// NewHealthAPIRouter creates a new mux router for the health api
func NewHealthAPIRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	alSvc := service.NewHealthService()
	alCtrl := controller.NewHealthAPIController(alSvc)
	router.HandleFunc("/alive", alCtrl.GetAliveMessage)
	return router, nil
}
