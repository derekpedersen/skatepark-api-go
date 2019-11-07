package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/derekpedersen/skatepark-api-go/model"
	log "github.com/sirupsen/logrus"
)

// SkateparkRepository interface
type SkateparkRepository interface {
	ParseSkateparks(dir string) ([]model.Skatepark, error)
}

// SkateparkRepositoryImpl implementation
type SkateparkRepositoryImpl struct {
}

// NewSkateparkRepository creates a new skate park repository
func NewSkateparkRepository() SkateparkRepository {
	return &SkateparkRepositoryImpl{}
}

// ParseSkateparks returns the collection of skateparks from a json file
// TODO: make this query a sql db, obvi
func (repo *SkateparkRepositoryImpl) ParseSkateparks(dir string) ([]model.Skatepark, error) {
	var result []model.Skatepark
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		raw, err := ioutil.ReadFile(dir + v.Name())
		if err != nil {
			log.Errorf("Error reading JSON file:\n %v", err)
			return nil, err
		}

		m := []model.Skatepark{}
		if err = json.Unmarshal(raw, &m); err != nil {
			log.Errorf("Error unmarshalling []model.Skatepark from %v: %v", dir+v.Name(), err)
			return nil, err
		}
		result = append(result, m...)
	}

	return result, nil
}
