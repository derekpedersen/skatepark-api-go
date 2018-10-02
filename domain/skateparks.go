package domain

import (
	"sort"

	"github.com/derekpedersen/skatepark-api-go/model"
)

// Skateparks represents the collection of skateparks
type Skateparks []model.Skatepark

func (sk Skateparks) Len() int           { return len(sk) }
func (sk Skateparks) Swap(i, j int)      { sk[i], sk[j] = sk[j], sk[i] }
func (sk Skateparks) Less(i, j int) bool { return sk[i].Address.State < sk[j].Address.State }

// GetSkateparksByCityMap organizes a map of skateparks by city
func (s Skateparks) GetSkateparksByCityMap() (cities map[string][]model.Skatepark) {
	cities = make(map[string][]model.Skatepark)
	for i := range s {
		if _, ok := cities[s[i].Address.City]; !ok {
			cities[s[i].Address.City] = []model.Skatepark{
				s[i],
			}
			continue
		}
		cities[s[i].Address.City] = append(cities[s[i].Address.City], s[i])
	}

	return
}

// GetSkateparksByStateMap oranizes a map of skateparks by state then by city
func (s Skateparks) GetSkateparksByStateMap() (states map[string]map[string][]model.Skatepark) {
	states = make(map[string]map[string][]model.Skatepark)

	// first go through and create a map of the states
	for i := range s {
		if _, ok := states[s[i].Address.State]; !ok {
			states[s[i].Address.State] = make(map[string][]model.Skatepark)
		}
	}

	// get a map of the cities
	cities := s.GetSkateparksByCityMap()

	// now loop through and get the cities within each state
	for k, v := range states {
		for i := range s {
			// determine if in the same state
			if k == s[i].Address.State {
				v[s[i].Address.City] = cities[s[i].Address.City]
			}
		}
	}

	return
}

// GetSkateparksByState gets the skateparks organized by state
func (s Skateparks) GetSkateparksByState() SkateparksByState {
	// get our handy map
	m := s.GetSkateparksByStateMap()
	states := make([]model.SkateparkByState, 0)

	// iterate over each state
	for km, vm := range m {
		state := model.SkateparkByState{
			State:  km,
			Cities: make([]model.SkateparkByCity, 0),
		}

		// iterate over each city
		for kc, vc := range vm {
			city := model.SkateparkByCity{
				City:       kc,
				Skateparks: vc,
			}

			state.Cities = append(state.Cities, city)
		}

		c := SkateparksByCity(state.Cities)
		sort.Sort(c)
		state.Cities = c

		states = append(states, state)
	}
	d := SkateparksByState(states)
	sort.Sort(d)
	return d
}
