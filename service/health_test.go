package service_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/service"
)

func TestHealthService_ReadyMessage(t *testing.T) {
	// Arrange
	svc := service.NewHealthService()

	// Act
	rm := svc.GetReadyMessage()

	// Assert
	if rm.Message != service.ReadyMessage {
		t.Errorf("expected a result of %v but was %v", service.ReadyMessage, rm.Message)
	}
}

func TestHealthService_AliveMessage(t *testing.T) {
	// Arrange
	svc := service.NewHealthService()

	// Act
	am := svc.GetAliveMessage()

	// Assert
	if am.Message != service.AliveMessage {
		t.Errorf("expected a result of %v but was %v", service.AliveMessage, am.Message)
	}
}

func TestHealthService_HealthyMessage(t *testing.T) {
	// Arrange
	svc := service.NewHealthService()

	// Act
	hm := svc.GetHealthyMessage()

	// Assert
	if hm.Message != service.HealthyMessage {
		t.Errorf("expected a result of %v but was %v", service.HealthyMessage, hm.Message)
	}
}
