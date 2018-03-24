package service

import (
	"encoding/json"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/model"
	"github.com/derekpedersen/skatepark-api-go/utils"
	"github.com/jeanphorn/log4go"
)

// GetSkateparks gets the full list of skateparks from json repository
func GetSkateparks() ([]model.Skatepark, error) {
	str, err := utils.ReadJsonFile("./repository/json/skateparks.json")
	if err != err {
		log4go.Error("Error reading JSON file:\n %v", err)
		return nil, err
	}
	res := []model.Skatepark{}
	if err = json.Unmarshal([]byte(str), &res); err != nil {
		log4go.Error("Error unmarshalling []Skatepark:\n %v", err)
		return nil, err
	}

	for i := range res {
		res[i].Album, err = imgurService.GetAlbum(res[i].AlbumID)
		if err != nil {
			log4go.Error("Error getting album:\n %v", err)
		}
	}

	return res, nil
}

// GetSkateparksByState gets a list of skateparks grouped by state from the json repository
func GetSkateparksByState() ([]model.SkateparkByState, error) {
	str, err := utils.ReadJsonFile("./skateparks/json/skateparks-states.json")
	if err != err {
		log4go.Error("Error reading JSON file:\n %v", err)
		return nil, err
	}

	res := []model.SkateparkByState{}
	if err = json.Unmarshal([]byte(str), &res); err != nil {
		log4go.Error("Error unmarshalling []SkateparkByState:\n %v", err)
		return nil, err
	}

	return res, nil
}
