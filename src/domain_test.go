package skatepark_api_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
)

func mockSkateparks() skatepark_api.Skateparks {
	return skatepark_api.Skateparks{
		{
			ID:          "1",
			Name:        "Liberty Skatepark",
			Description: "A great skatepark in Liberty City.",
			Address: skatepark_api.Address{
				AddressLine1: "123 Main St",
				City:         "Liberty City",
				State:        "New York",
				ZipCode:      "10001",
			},
		},
		{
			ID:          "2",
			Name:        "Downtown Ramps",
			Description: "Street style ramps downtown.",
			Address: skatepark_api.Address{
				AddressLine1: "456 Side St",
				City:         "Liberty City",
				State:        "New York",
				ZipCode:      "10002",
			},
		},
		{
			ID:          "3",
			Name:        "Hilltop Skate Zone",
			Description: "Nice views from the top.",
			Address: skatepark_api.Address{
				AddressLine1: "789 Hilltop Rd",
				City:         "Springfield",
				State:        "Illinois",
				ZipCode:      "62704",
			},
		},
	}
}

func Test_GetSkateparkByName(t *testing.T) {

	// Arrange
	skateparks := mockSkateparks()

	t.Run("Exact match", func(t *testing.T) {
		// Act
		result := skateparks.GetSkateparkByName("Liberty Skatepark")
		// Assert
		assert.Equal(t, "Liberty Skatepark", result.Name)
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		// Act
		result := skateparks.GetSkateparkByName("downtown ramps")
		// Assert
		assert.Equal(t, "Downtown Ramps", result.Name)
	})

	t.Run("No match returns empty", func(t *testing.T) {
		// Act
		result := skateparks.GetSkateparkByName("Not Real Park")
		// Assert
		assert.True(t, strings.TrimSpace(result.ID) == "")
	})
}

func Test_CitySkateparkMap(t *testing.T) {

	// Arrange
	skateparks := mockSkateparks()

	// Act
	cityMap := skateparks.CitySkateparkMap()

	// Assert
	assert.Len(t, cityMap, 2)

	assert.Contains(t, cityMap, "Liberty City")
	assert.Contains(t, cityMap, "Springfield")

	assert.Len(t, cityMap["Liberty City"], 2)
	assert.Len(t, cityMap["Springfield"], 1)
}

func Test_StateSkateparkMap(t *testing.T) {

	// Arrange
	skateparks := mockSkateparks()

	// Act
	stateMap := skateparks.StateSkateparkMap()

	// Assert
	assert.Len(t, stateMap, 2)

	assert.Contains(t, stateMap, "New York")
	assert.Contains(t, stateMap, "Illinois")

	assert.Contains(t, stateMap["New York"], "Liberty City")
	assert.Contains(t, stateMap["Illinois"], "Springfield")

	assert.Len(t, stateMap["New York"]["Liberty City"], 2)
	assert.Len(t, stateMap["Illinois"]["Springfield"], 1)
}
