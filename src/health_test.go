package skatepark_api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
	"github.com/stretchr/testify/assert"
)

func Test_HealthServiceImpl(t *testing.T) {
	svc := skatepark_api.NewHealthService()

	tests := []struct {
		name     string
		method   func() *skatepark_api.HealthMessage
		expected string
	}{
		{
			name:     "GetHealthyMessage returns correct message",
			method:   svc.GetHealthyMessage,
			expected: skatepark_api.HealthyMessage,
		},
		{
			name:     "GetAliveMessage returns correct message",
			method:   svc.GetAliveMessage,
			expected: skatepark_api.AliveMessage,
		},
		{
			name:     "GetReadyMessage returns correct message",
			method:   svc.GetReadyMessage,
			expected: skatepark_api.ReadyMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method()
			assert.Equal(t, tt.expected, result.Message)
		})
	}
}

func Test_HealthAPIControllerImpl(t *testing.T) {
	svc := skatepark_api.NewHealthService()
	api := skatepark_api.NewHealthAPIController(svc)

	tests := []struct {
		name           string
		handler        func(http.ResponseWriter, *http.Request)
		expectedStatus int
		expectedBody   *skatepark_api.HealthMessage
	}{
		{
			name:           "GetHealthyMessage returns 200 and correct message",
			handler:        api.GetHealthyMessage,
			expectedStatus: http.StatusOK,
			expectedBody:   &skatepark_api.HealthMessage{Message: skatepark_api.HealthyMessage},
		},
		{
			name:           "GetAliveMessage returns 200 and correct message",
			handler:        api.GetAliveMessage,
			expectedStatus: http.StatusOK,
			expectedBody:   &skatepark_api.HealthMessage{Message: skatepark_api.AliveMessage},
		},
		{
			name:           "GetReadyMessage returns 200 and correct message",
			handler:        api.GetReadyMessage,
			expectedStatus: http.StatusOK,
			expectedBody:   &skatepark_api.HealthMessage{Message: skatepark_api.ReadyMessage},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			tt.handler(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var result skatepark_api.HealthMessage
			err := json.Unmarshal(rr.Body.Bytes(), &result)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody.Message, result.Message)
		})
	}
}
