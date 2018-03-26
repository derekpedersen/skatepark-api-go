package repository

import (
	"encoding/json"

	"github.com/derekpedersen/skatepark-api-go/model"
	"github.com/derekpedersen/skatepark-api-go/utils"
	"github.com/jeanphorn/log4go"
)

// GetSkateparks returns the collection of skateparks
func GetSkateparks() (m []model.Skatepark, err error) {
	str, err := utils.ReadJsonFile("./repository/json/skateparks.json")

	if err != err {
		log4go.Error("Error reading JSON file:\n %v", err)
		return nil, err
	}

	if err = json.Unmarshal([]byte(str), &m); err != nil {
		log4go.Error("Error unmarshalling []Skatepark:\n %v", err)
		return nil, err
	}

	return m, nil
}
