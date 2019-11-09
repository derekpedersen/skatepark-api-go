package appcfg_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/mock"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/golang/mock/gomock"

	"github.com/derekpedersen/skatepark-api-go/appcfg"
)

func TestNewBaseRouter(t *testing.T) {
	// Arrange

	// Act
	r, err := appcfg.NewBaseRouter()

	// Assert
	if err != nil {
		t.Errorf("encounterd an unexpecterd error: %v", err)
	}
	if r == nil {
		t.Error("expected router to be set")
	}
}

func TestAddHealthRoutes(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := service.NewHealthService()
	api := controller.NewHealthAPIController(svc)
	router, err := appcfg.NewBaseRouter()
	if err != nil {
		t.Fatalf("encountered unexpected error: %v", err)
	}

	// Act
	appcfg.AddHealthRoutes(router, api)

	// Assert
	// ??
}

func TestNewSkateparkAPIRouter(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	api := mock.NewMockSkateparksAPIController(ctrl)
	router, err := appcfg.NewBaseRouter()
	if err != nil {
		t.Fatalf("encountered an unexpected error: %v", err)
	}

	// Act
	r, err := appcfg.NewSkateparkAPIRouter(router, api)

	// Assert
	if err != nil {
		t.Fatalf("encounterd an unexpected error: %v", err)
	}
	if r == nil {
		t.Error("expected the router to be set")
	}
}
