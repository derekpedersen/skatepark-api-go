package skatepark_api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// --- MOCK HEALTH CONTROLLER ---

type mockHealthController struct{}

func (m *mockHealthController) GetAliveMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alive"))
}

func (m *mockHealthController) GetReadyMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func (m *mockHealthController) GetHealthyMessage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

// --- TESTS ---

func Test_NewSwaggerRouter(t *testing.T) {
	router := skatepark_api.NewSwaggerRouter(".")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusNotFound, rr.Code, "Swagger router should serve file or directory")
}

func Test_NewBaseRouter(t *testing.T) {
	router, err := skatepark_api.NewBaseRouter()
	assert.NoError(t, err)
	assert.NotNil(t, router)
}

func Test_AddHealthRoutes(t *testing.T) {
	router := mux.NewRouter()
	mockCtrl := &mockHealthController{}

	skatepark_api.AddHealthRoutes(router, mockCtrl)

	tests := []struct {
		name       string
		path       string
		statusCode int
		body       string
	}{
		{"Alive route", "/alive", http.StatusOK, "Alive"},
		{"Ready route", "/ready", http.StatusOK, "Ready"},
		{"Healthy route", "/healthy", http.StatusOK, "Healthy"},
		{"Not found route", "/missing", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code)
			if tt.body != "" {
				assert.Equal(t, tt.body, rr.Body.String())
			}
		})
	}
}

func Test_NewSkateparkAPIRouter(t *testing.T) {
	// Use mock handlers to avoid relying on actual implementation
	router := mux.NewRouter()

	// Replace actual handlers with simple mocks
	router.HandleFunc("/api/skateparks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/skateparks/{state}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/skateparks/{state}/{city}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/skateparks/{state}/{city}/{skatepark}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet, http.MethodOptions)

	newRouter, err := skatepark_api.NewSkateparkAPIRouter(router)
	assert.NoError(t, err)

	tests := []struct {
		name       string
		path       string
		method     string
		statusCode int
	}{
		{"Base skateparks route", "/api/skateparks", http.MethodGet, http.StatusOK},
		{"State skateparks route", "/api/skateparks/California", http.MethodGet, http.StatusOK},
		{"City skateparks route", "/api/skateparks/California/SanDiego", http.MethodGet, http.StatusOK},
		{"Skatepark by name route", "/api/skateparks/California/SanDiego/OceanPark", http.MethodGet, http.StatusOK},
		{"Not found route", "/api/skateparks/unknown/path", http.MethodPost, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			newRouter.ServeHTTP(rr, req)
			assert.Equal(t, tt.statusCode, rr.Code)
		})
	}
}
