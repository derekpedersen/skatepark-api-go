package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/derekpedersen/skatepark-api-go/domain"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/mock"
	"github.com/derekpedersen/skatepark-api-go/model"
	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
)

func TestNewSkateparksAPIController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockSkateparksService(ctrl)

	// Act
	r := controller.NewSkateparksAPIController(svc)

	// Assert
	if r == nil {
		t.Error("expected controller to be set")
	}
}

func TestGetSkateparks(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockSkateparksService(ctrl)
	api := controller.NewSkateparksAPIController(svc)

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
				url:    "/skateparks/states",
			},
			statuscode: 200,
		},
	}

	skateparks := domain.Skateparks(
		[]model.Skatepark{},
	)
	gomock.InOrder(
		// GetHealthyMessage: 200
		svc.EXPECT().GetSkateparks().Return(skateparks, nil),
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
			handler := http.HandlerFunc(api.GetSkateparksByState)
			handler.ServeHTTP(rr, req)

			// Assert
			if rr.Code != tt.statuscode {
				t.Errorf("unexpected status code: have %v, want %v", rr.Code, tt.statuscode)
			}
		})
	}
}
