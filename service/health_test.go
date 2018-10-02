package service

import "testing"

func TestHealthService(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{
			Name: "Constructor",
		},
	}

	// Arrange
	svc := NewHealthService()

	for _, test := range tests {
		// Act
		rm := svc.GetReadyMessage() // TODO: break these out into separate tests
		am := svc.GetAliveMessage()

		// Assert
		if rm.Message != "Johnny 5 Ready!" {
			t.Errorf("%v: expected a result of %v but was %v", "Johnny 5 Ready!", test.Name, rm)
		}
		if am.Message != "Johnny 5 Alive!" {
			t.Errorf("%v: expected a result of %v but was %v", "Johnny 5 Alive!", test.Name, am)
		}
	}
}
