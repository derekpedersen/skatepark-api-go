package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/domain"
	"github.com/derekpedersen/skatepark-api-go/mock"
	"github.com/derekpedersen/skatepark-api-go/model"

	"github.com/golang/mock/gomock"
)

func TestNewHealthAPIController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	skateSvc := mock.NewMockSkateparksService(ctrl)

	// Act
	r := controller.NewHealthAPIController(svc, skateSvc)

	// Assert
	if r == nil {
		t.Error("expected controller to be set")
	}
}

func TestGetAliveMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	skateSvc := mock.NewMockSkateparksService(ctrl)
	api := controller.NewHealthAPIController(svc, skateSvc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetAliveMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/alive",
			},
			statuscode: 200,
		},
	}

	msg := model.HealthMessage{
		Message: "Howdy",
	}

	skateparks := domain.Skateparks(
		[]model.Skatepark{},
	)

	gomock.InOrder(
		// GetAliveMessage: 200
		svc.EXPECT().GetAliveMessage().Return(&msg),
		skateSvc.EXPECT().GetSkateparks().Return(skateparks, nil),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetAliveMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}

func TestGetReadyMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	skateSvc := mock.NewMockSkateparksService(ctrl)
	api := controller.NewHealthAPIController(svc, skateSvc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetReadyMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/ready",
			},
			statuscode: 200,
		},
	}

	msg := model.HealthMessage{
		Message: "Howdy",
	}

	skateparks := domain.Skateparks(
		[]model.Skatepark{},
	)

	gomock.InOrder(
		// GetAliveMessage: 200
		svc.EXPECT().GetReadyMessage().Return(&msg),
		skateSvc.EXPECT().GetSkateparks().Return(skateparks, nil),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetReadyMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}

func TestGetHealthyMessage(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)
	skateSvc := mock.NewMockSkateparksService(ctrl)
	api := controller.NewHealthAPIController(svc, skateSvc)

	type args struct {
		method string
		url    string
		params map[string]string
	}
	tests := []struct {
		name       string
		args       args
		statuscode int
	}{
		{
			name: "GetHealthyMessage: 200",
			args: args{
				method: http.MethodGet,
				url:    "/healthy",
			},
			statuscode: 200,
		},
	}

	msg := model.HealthMessage{
		Message: "Howdy",
	}

	skateparks := domain.Skateparks(
		[]model.Skatepark{},
	)

	gomock.InOrder(
		// GetHealthyMessage: 200
		svc.EXPECT().GetHealthyMessage().Return(&msg),
		skateSvc.EXPECT().GetSkateparks().Return(skateparks, nil),
	)

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, tt.args.url, nil)
			if err != nil {
				t.Fatalf("encountered an unexpected error: %v", err)
			}

			rr := httptest.NewRecorder()
			req = mux.SetURLVars(req, tt.args.params)
			handler := http.HandlerFunc(api.GetHealthyMessage)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}
