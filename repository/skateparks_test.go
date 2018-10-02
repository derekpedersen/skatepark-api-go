package repository

import (
	"testing"
)

func TestSkateparks(t *testing.T) {
	tests := []struct {
		Name string
	}{
		{
			Name: "Constructor",
		},
	}
	// Arrange
	repo := NewSkateparkRepository()

	for _, test := range tests {
		// Act
		parks, err := repo.ParseSkateparks("./json/skateparks.json")

		// Assert
		if err != nil {
			t.Fatalf("%v, encountered an unexpected error: %v", test.Name, err)
		}
		if len(parks) == 0 {
			t.Errorf("%v: expected a result length greater than 0", test.Name)
		}
	}
}
