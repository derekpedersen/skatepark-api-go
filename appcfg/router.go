package appcfg

import (
	"net/http"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/gorilla/mux"
)

// NewBaseRouter creates a new base router with standard health checks
func NewBaseRouter() (*mux.Router, error) {
	return mux.NewRouter(), nil
}

// AddHealthRoutes adds alive, ready, and healthy routes to the router
func AddHealthRoutes(
	router *mux.Router,
	healthctrl controller.HealthAPIController,
) {
	router.HandleFunc("/alive", healthctrl.GetAliveMessage)
	router.HandleFunc("/ready", healthctrl.GetReadyMessage)
	router.HandleFunc("/healthy", healthctrl.GetHealthyMessage)
}

// NewSkateparkAPIRouter creates a new mux router for the skatepark api
func NewSkateparkAPIRouter(
	router *mux.Router,
	imgur imgurService.AlbumService,
	skatectrl controller.SkateparksAPIController,
) (*mux.Router, error) {
	// api subrouter
	api := router.StrictSlash(true).PathPrefix("/api").Subrouter()

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: true
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	//api.HandleFunc("/skateparks", skatectrl.GetSkateparks).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: true
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	api.HandleFunc("/skatepark/states", skatectrl.GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: true
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	api.HandleFunc("/skatepark/{state}", skatectrl.GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: true
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	api.HandleFunc("/skatepark/{state}/{city}", skatectrl.GetSkateparksByCity).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Deprecated: true
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	api.HandleFunc("/skatepark/{state}/{city}/{skatepark}", skatectrl.GetSkateparksByName).Methods(http.MethodGet, http.MethodOptions)
	return router, nil
}
