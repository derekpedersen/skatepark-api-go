package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/derekpedersen/skatepark-api-go/model"
	log "github.com/sirupsen/logrus"
)

// SkateparkRepository interface
type SkateparkRepository interface {
	ParseSkateparks(filepath string) ([]model.Skatepark, error)
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
func (repo *SkateparkRepositoryImpl) ParseSkateparks(filepath string) ([]model.Skatepark, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != err {
		log.Errorf("Error reading JSON file:\n %v", err)
		return nil, err
	}

	str := fmt.Sprintf("%s", raw)

	m := []model.Skatepark{}
	if err = json.Unmarshal([]byte(str), &m); err != nil {
		log.Errorf("Error unmarshalling []model.Skatepark:\n %v", err)
		return nil, err
	}

	return m, nil
}
