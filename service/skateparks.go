package service

import (
	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/model"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/jeanphorn/log4go"
)

func init() {
	imgurService.NewAlbumService("ae20e9f5a365875")
}

// GetSkateparks gets the full list of skateparks from json repository
func GetSkateparks() (m []model.Skatepark, err error) {
	m, err = repository.GetSkateparks()

	for i := range m {
		m[i].Album, err = imgurService.GetAlbum(m[i].AlbumID)
		if err != nil {
			log4go.Error("Error getting album:\n %v", err)
		}
	}

	return m, nil
}

// GetSkateparksByState gets a list of skateparks grouped by state from the json repository
func GetSkateparksByState() (m []model.SkateparkByState, err error) {
	skateparks, err := GetSkateparks()
	if err != err {
		log4go.Error("Error getting skateparks", err)
		return nil, err
	}

	states := make(map[string][]model.Skatepark)
	for i := range skateparks {
		if _, ok := states[skateparks[i].Address.State]; !ok {
			states[skateparks[i].Address.State] = []model.Skatepark{}
		}

		states[skateparks[i].Address.State] = append(states[skateparks[i].Address.State], skateparks[i])
	}

	for k, v := range states {
		s := model.SkateparkByState{
			State: k,
		}

		cities := make(map[string][]model.Skatepark)

		for i := range v {
			if _, ok := cities[v[i].Address.City]; !ok {
				cities[v[i].Address.City] = []model.Skatepark{}
			}

			cities[v[i].Address.City] = append(cities[v[i].Address.City], v[i])
		}

		for ky, vl := range cities {
			sc := model.SkateparkByCity{
				City:       ky,
				Skateparks: vl,
			}
			s.Cities = append(s.Cities, sc)
		}

		m = append(m, s)
	}

	return m, nil
}
