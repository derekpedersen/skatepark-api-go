package domain

import (
	"strings"

	"github.com/derekpedersen/skatepark-api-go/model"
)

// Skateparks represents the collection of skateparks
type Skateparks []model.Skatepark

func (dom Skateparks) Len() int           { return len(dom) }
func (dom Skateparks) Swap(i, j int)      { dom[i], dom[j] = dom[j], dom[i] }
func (dom Skateparks) Less(i, j int) bool { return dom[i].Address.State < dom[j].Address.State }

// CitySkateparkMap is a map of skateparks with the city as a key
// swagger:model
type CitySkateparkMap map[string]Skateparks

// StateSkateparkMap is a map of skateparks with the state as a key
// swagger:model
type StateSkateparkMap map[string]CitySkateparkMap

// GetSkateparkByName does just that!
func (dom Skateparks) GetSkateparkByName(name string) model.Skatepark {
	s := model.Skatepark{}
	for i := range dom {
		if strings.ToLower(dom[i].Name) == strings.ToLower(name) {
			s = dom[i]
		}
	}
	return s
}

// CitySkateparkMap organizes a map of skateparks by city
func (dom Skateparks) CitySkateparkMap() CitySkateparkMap {
	cities := make(CitySkateparkMap)
	for i := range dom {
		if _, ok := cities[dom[i].Address.City]; !ok {
			cities[dom[i].Address.City] = []model.Skatepark{
				dom[i],
			}
			continue
		}
		cities[dom[i].Address.City] = append(cities[dom[i].Address.City], dom[i])
	}

	return cities
}

// StateSkateparkMap oranizes a map of skateparks by state then by city
func (dom Skateparks) StateSkateparkMap() StateSkateparkMap {
	states := make(StateSkateparkMap)

	// first go through and create a map of the states
	for i := range dom {
		if _, ok := states[dom[i].Address.State]; !ok {
			states[dom[i].Address.State] = make(CitySkateparkMap)
		}
	}

	// get a map of the cities
	cities := dom.CitySkateparkMap()

	// now loop through and get the cities within each state
	for k, v := range states {
		for i := range dom {
			// determine if in the same state
			if k == dom[i].Address.State {
				v[dom[i].Address.City] = cities[dom[i].Address.City]
			}
		}
	}

	return states
}
