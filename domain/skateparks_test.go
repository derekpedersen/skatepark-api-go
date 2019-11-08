package domain_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/domain"

	"github.com/derekpedersen/skatepark-api-go/model"
)

func TestSkateparks(t *testing.T) {
	// Arrange
	skateparks := []model.Skatepark{
		model.Skatepark{
			Name: "ted",
		},
		model.Skatepark{
			Name: "bill",
		},
	}
	dom := domain.Skateparks(skateparks)

	// Act
	r := dom.GetSkateparkByName("ted")

	// Assert
	if r.Name != "ted" {
		t.Error("expected to find a result")
	}
}
