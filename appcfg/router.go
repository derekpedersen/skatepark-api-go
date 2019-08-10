package appcfg

import (
	"net/http"
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

	router := mux.NewRouter()

	// health routes
	hsvc := service.NewHealthService()
	hctrl := controller.NewHealthAPIController(hsvc)
	router.HandleFunc("/alive", hctrl.GetAliveMessage)
	router.HandleFunc("/ready", hctrl.GetReadyMessage)
	router.HandleFunc("/healthy", hctrl.GetHealthyMessage)

	// api subrouter
	api := router.StrictSlash(true).PathPrefix("/api").Subrouter()

	// skatepark routes
	skRepo := repository.NewSkateparkRepository()
	skSvc := service.NewSkateparksService(imgSvc, skRepo)
	skCtrl := controller.NewSkateparksAPIController(skSvc)
	//api.HandleFunc("/skateparks", skCtrl.GetSkateparks).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/states", skCtrl.GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/states/{state}", skCtrl.GetSkateparksByCity).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/states/{state}/{city}", skCtrl.GetSkateparksByCity).Methods(http.MethodGet, http.MethodOptions)
	return router, nil
}
