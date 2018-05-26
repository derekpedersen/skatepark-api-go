package appcfg

import (
	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/gorilla/mux"
)

func NewSkateparkAPIRouter() (*mux.Router, error) {
	// set "api" path prefix
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()

	skRepo := repository.NewSkateparkRepository()
	skSvc := service.NewSkateparksService(skRepo)
	skCtrl := controller.NewSkateparksAPIController(skSvc)
	router.HandleFunc("/skateparks", skCtrl.GetSkateparks)
	router.HandleFunc("/skatepark/states", skCtrl.GetSkateparksByState) // TODO: this is the legacy one, but need to update SPA first before removing

	return router, nil
}

func NewAliveAPIRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	alSvc := service.NewAliveService()
	alCtrl := controller.NewAliveAPIController(alSvc)
	router.HandleFunc("/alive", alCtrl.GetAliveMessage)
	return router, nil
}
