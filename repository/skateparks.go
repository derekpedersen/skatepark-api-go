package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			raw, err := ioutil.ReadFile(path)
			if err != nil {
				log.Errorf("Error reading JSON file:\n %v", err)
				return err
			}

			m := []model.Skatepark{}
			if err = json.Unmarshal(raw, &m); err != nil {
				log.Errorf("Error unmarshalling []model.Skatepark from %v: %v", dir+info.Name(), err)
				return err
			}
			result = append(result, m...)

			return nil
		})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return result, nil
}
