package controller_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/mock"

	"github.com/golang/mock/gomock"
)

func TestNewHealthAPIController(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	svc := mock.NewMockHealthService(ctrl)

	// Act
	r := controller.NewHealthAPIController(svc)

	// Assert
	if r == nil {
		t.Error("expected controller to be set")
	}
}
