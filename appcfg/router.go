package appcfg

import (
	"net/http"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/gorilla/mux"
)

// NewBaseRouter creates a new base router with standard health checks
func NewBaseRouter() (*mux.Router, error) {
	router := mux.NewRouter()

	// health routes
	hsvc := service.NewHealthService()
	hctrl := controller.NewHealthAPIController(hsvc)
	router.HandleFunc("/alive", hctrl.GetAliveMessage)
	router.HandleFunc("/ready", hctrl.GetReadyMessage)
	router.HandleFunc("/healthy", hctrl.GetHealthyMessage)

	return router, nil
}

// NewSkateparkAPIRouter creates a new mux router for the skatepark api
func NewSkateparkAPIRouter(
	router *mux.Router,
	imgur imgurService.AlbumService,
	skatectrl controller.SkateparksAPIController,
) (*mux.Router, error) {
	// api subrouter
	api := router.StrictSlash(true).PathPrefix("/api").Subrouter()

	//api.HandleFunc("/skateparks", skatectrl.GetSkateparks).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/states", skatectrl.GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/{state}", skatectrl.GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/{state}/{city}", skatectrl.GetSkateparksByCity).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/skatepark/{state}/{city}/{skatepark}", skatectrl.GetSkateparksByName).Methods(http.MethodGet, http.MethodOptions)
	return router, nil
}
