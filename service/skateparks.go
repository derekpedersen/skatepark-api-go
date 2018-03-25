package service

import (
	"os"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/domain"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/jeanphorn/log4go"
)

// SkateparksService interface
type SkateparksService interface {
	GetSkateparks() (domain.Skateparks, error)
}

// SkateparksServiceImpl implementation
type SkateparksServiceImpl struct {
	repo   repository.SkateparkRepository
	imgSvc imgurService.AlbumsService
}

// NewSkateparksService creates a new skateparks service
func NewSkateparksService(repo repository.SkateparkRepository) *SkateparksServiceImpl {
	apiKey := os.Getenv("IMGUR_API_KEY")
	imgSvc := imgurService.NewAlbumService(apiKey)
	return &SkateparksServiceImpl{
		repo:   repo,
		imgSvc: imgSvc,
	}
}

// GetSkateparks gets the full list of skateparks from json repository
func (svc *SkateparksServiceImpl) GetSkateparks() (domain.Skateparks, error) {
	m, err := svc.repo.GetSkateparks()
	if err != nil {
		return nil, err
	}

	for i := range m {
		m[i].Album, err = svc.imgSvc.GetAlbum(m[i].AlbumID)
		if err != nil {
			log4go.Error("Error getting album:\n %v", err)
		}
	}

	d := domain.Skateparks(m)
	return d, nil
}
