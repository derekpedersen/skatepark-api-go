package controller_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/controller"
	"github.com/derekpedersen/skatepark-api-go/mock"

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
