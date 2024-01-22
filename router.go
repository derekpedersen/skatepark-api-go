package skatepark_api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewSwaggerRouter creates a new swagger router
func NewSwaggerRouter(directory string) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(directory)))
	return r
}

// NewBaseRouter creates a new base router with standard health checks
func NewBaseRouter() (*mux.Router, error) {
	return mux.NewRouter(), nil
}

// AddHealthRoutes adds alive, ready, and healthy routes to the router
func AddHealthRoutes(
	router *mux.Router,
	healthctrl HealthAPIController,
) {
	router.HandleFunc("/alive", healthctrl.GetAliveMessage)
	router.HandleFunc("/ready", healthctrl.GetReadyMessage)
	router.HandleFunc("/healthy", healthctrl.GetHealthyMessage)
}

// NewSkateparkAPIRouter creates a new mux router for the skateparks api
func NewSkateparkAPIRouter(
	router *mux.Router,
) (*mux.Router, error) {
	// api subrouter
	api := router.StrictSlash(true).PathPrefix("/api").Subrouter()

	// swagger:route GET /skateparks skateparks listSkateparks
	//
	// Lists skateparks filtered by some parameters.
	//
	// This will show all available skateparks by default.
	// You can get the skateparks that are out of stock
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
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: []skatepark
	api.HandleFunc("/skateparks", GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /skateparks/:state stateSkatepArks listStateSkateparks
	//
	// Lists skateparks filtered by some parameters.
	//
	// This will show all available skateparks by default.
	// You can get the skateparks that are out of stock
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
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: stateSkateparkMap
	api.HandleFunc("/skateparks/{state}", GetSkateparksByState).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /skateparks/:state/:city citySkateparks listCitySkateparks
	//
	// Lists skateparks filtered by some parameters.
	//
	// This will show all available skateparks by default.
	// You can get the skateparks that are out of stock
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
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: citySkateparkMap
	api.HandleFunc("/skateparks/{state}/{city}", GetSkateparksByCity).Methods(http.MethodGet, http.MethodOptions)

	// swagger:route GET /skateparks/:state/:city/:skatepark skateparks getSkatepark
	//
	// Lists skateparks filtered by some parameters.
	//
	// This will show all available skateparks by default.
	// You can get the skateparks that are out of stock
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
	//       200: skatepark
	api.HandleFunc("/skateparks/{state}/{city}/{skatepark}", GetSkateparksByName).Methods(http.MethodGet, http.MethodOptions)
	return router, nil
}
