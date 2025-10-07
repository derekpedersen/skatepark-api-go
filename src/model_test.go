package skatepark_api_test

import (
	"encoding/json"
	"testing"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
	"github.com/stretchr/testify/assert"
)

// Mocks
var mockGoogleMaps = skatepark_api.GoogleMaps{
	EmbedMap:  "https://maps.google.com/embed/mock",
	ShareLink: "https://maps.google.com/share/mock",
}

var mockAddress = skatepark_api.Address{
	AddressLine1: "123 Skate Lane",
	AddressLine2: "Suite 100",
	City:         "Kickflip",
	State:        "Ollinois",
	ZipCode:      "12345",
	GoogleMaps:   mockGoogleMaps,
}

var mockSkatepark = skatepark_api.Skatepark{
	ID:          "1",
	Name:        "Tony Hawk's Paradise",
	Description: "Legendary skatepark with a half-pipe.",
	Address:     mockAddress,
	AlbumID:     "album123",
	Album:       nil, // Assume album tests are separate
}

func TestGoogleMapsJSON(t *testing.T) {
	input := mockGoogleMaps

	t.Run("Marshal and Unmarshal GoogleMaps", func(t *testing.T) {
		data, err := json.Marshal(input)
		assert.NoError(t, err)

		var result skatepark_api.GoogleMaps
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)
		assert.Equal(t, input, result)
	})
}

func TestAddressJSON(t *testing.T) {
	input := mockAddress

	t.Run("Marshal and Unmarshal Address", func(t *testing.T) {
		data, err := json.Marshal(input)
		assert.NoError(t, err)

		var result skatepark_api.Address
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)
		assert.Equal(t, input, result)
	})
}

func TestSkateparkJSON(t *testing.T) {
	input := mockSkatepark

	t.Run("Marshal and Unmarshal Skatepark", func(t *testing.T) {
		data, err := json.Marshal(input)
		assert.NoError(t, err)

		var result skatepark_api.Skatepark
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)

		assert.Equal(t, input.ID, result.ID)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Address.City, result.Address.City)
		assert.Equal(t, input.Address.GoogleMaps.EmbedMap, result.Address.GoogleMaps.EmbedMap)
	})
}

func TestSkateparks_SortingMethods(t *testing.T) {
	parks := skatepark_api.Skateparks{
		{
			ID:      "1",
			Name:    "Park A",
			Address: skatepark_api.Address{State: "California"},
		},
		{
			ID:      "2",
			Name:    "Park B",
			Address: skatepark_api.Address{State: "Arizona"},
		},
	}

	t.Run("Len returns correct length", func(t *testing.T) {
		assert.Equal(t, 2, parks.Len())
	})

	t.Run("Swap switches values", func(t *testing.T) {
		parks.Swap(0, 1)
		assert.Equal(t, "Park B", parks[0].Name)
		assert.Equal(t, "Park A", parks[1].Name)
	})

	t.Run("Less returns true when i < j by state", func(t *testing.T) {
		parks = skatepark_api.Skateparks{
			{Address: skatepark_api.Address{State: "Alaska"}},
			{Address: skatepark_api.Address{State: "Texas"}},
		}
		assert.True(t, parks.Less(0, 1))
		assert.False(t, parks.Less(1, 0))
	})
}
