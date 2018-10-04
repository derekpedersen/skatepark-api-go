package service

import "testing"

func TestHealthService_ReadyMessage(t *testing.T) {
	// Arrange
	svc := NewHealthService()

	// Act
	rm := svc.GetReadyMessage()

	// Assert
	if rm.Message != ReadyMessage {
		t.Errorf("expected a result of %v but was %v", ReadyMessage, rm.Message)
	}
}

func TestHealthService_AliveMessage(t *testing.T) {
	// Arrange
	svc := NewHealthService()

	// Act
	am := svc.GetAliveMessage()

	// Assert
	if am.Message != AliveMessage {
		t.Errorf("expected a result of %v but was %v", AliveMessage, am.Message)
	}
}

func TestHealthService_HealthyMessage(t *testing.T) {
	// Arrange
	svc := NewHealthService()

	// Act
	hm := svc.GetHealthyMessage()

	// Assert
	if hm.Message != HealthyMessage {
		t.Errorf("expected a result of %v but was %v", HealthyMessage, hm.Message)
	}
}
