package domain_test

import (
	"testing"

	"github.com/derekpedersen/skatepark-api-go/domain"

	"github.com/derekpedersen/skatepark-api-go/model"
)

func TestGetSkateparkByName(t *testing.T) {
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

func TestCitySkateparkMap(t *testing.T) {
	// Arrange
	skateparks := []model.Skatepark{
		model.Skatepark{
			Name: "ted",
			Address: model.Address{
				City: "ted",
			},
		},
		model.Skatepark{
			Name: "bill",
			Address: model.Address{
				City: "bill",
			},
		},
		model.Skatepark{
			Name: "derek",
			Address: model.Address{
				City: "bill",
			},
		},
	}
	dom := domain.Skateparks(skateparks)

	// Act
	r := dom.CitySkateparkMap()

	// Assert
	if len(r) != 2 {
		t.Error("expected a result of 2 cities")
	}
}

func TestStateSkateparkMap(t *testing.T) {
	// Arrange
	skateparks := []model.Skatepark{
		model.Skatepark{
			Name: "ted",
			Address: model.Address{
				City:  "ted",
				State: "WA",
			},
		},
		model.Skatepark{
			Name: "bill",
			Address: model.Address{
				City:  "bill",
				State: "WA",
			},
		},
	}
	dom := domain.Skateparks(skateparks)

	// Act
	r := dom.StateSkateparkMap()

	// Assert
	if len(r) != 1 {
		t.Error("expected a result of 1 state")
	}
}
