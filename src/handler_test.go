package skatepark_api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
)

// Override SKATEPARKS with test data
var mockData = skatepark_api.Skateparks{
	{
		ID:   "1",
		Name: "Sunset Park",
		Address: skatepark_api.Address{
			City:  "San Diego",
			State: "California",
		},
	},
	{
		ID:   "2",
		Name: "Ocean Skate",
		Address: skatepark_api.Address{
			City:  "San Diego",
			State: "California",
		},
	},
	{
		ID:   "3",
		Name: "Windy Skatepark",
		Address: skatepark_api.Address{
			City:  "Chicago",
			State: "Illinois",
		},
	},
}

func setup() {
	skatepark_api.SKATEPARKS = mockData
}

func Test_GetSkateparksByState(t *testing.T) {
	setup()

	t.Run("Valid state returns skateparks", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/state/California", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/state/{state}", skatepark_api.GetSkateparksByState)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var result skatepark_api.CitySkateparkMap
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Contains(t, result, "San Diego")
		assert.Len(t, result["San Diego"], 2)
	})

	t.Run("Unknown state returns 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/state/Texas", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/state/{state}", skatepark_api.GetSkateparksByState)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func Test_GetSkateparksByCity(t *testing.T) {
	setup()

	t.Run("Valid city returns skateparks", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/city/San%20Diego", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/city/{city}", skatepark_api.GetSkateparksByCity)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var result []skatepark_api.Skatepark
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("Unknown city returns 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/city/New%20York", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/city/{city}", skatepark_api.GetSkateparksByCity)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func Test_GetSkateparksByName(t *testing.T) {
	setup()

	t.Run("Valid name returns skatepark", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/name/Ocean%20Skate", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/name/{skatepark}", skatepark_api.GetSkateparksByName)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var result skatepark_api.Skatepark
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Equal(t, "Ocean Skate", result.Name)
	})

	t.Run("Unknown name returns empty object", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/skateparks/name/UnknownPark", nil)
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/skateparks/name/{skatepark}", skatepark_api.GetSkateparksByName)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var result skatepark_api.Skatepark
		err := json.Unmarshal(rr.Body.Bytes(), &result)
		assert.NoError(t, err)
		assert.Empty(t, result.ID)
	})
}
