package appcfg_test

import (
	"testing"

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
