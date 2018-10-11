package service

import (
	"os"
	"time"

	imgurService "github.com/derekpedersen/imgur-go/service"
	"github.com/derekpedersen/skatepark-api-go/domain"
	"github.com/derekpedersen/skatepark-api-go/repository"
	log "github.com/sirupsen/logrus"
)

// SkateparksService interface
type SkateparksService interface {
	GetSkateparks() (domain.Skateparks, error)
}

// SkateparksServiceImpl implementation
type SkateparksServiceImpl struct {
	repo       repository.SkateparkRepository
	imgSvc     imgurService.AlbumService
	parkscache struct {
		skateparks  domain.Skateparks
		refreshedat time.Time
	}
}

// NewSkateparksService creates a new skateparks service
func NewSkateparksService(imgSvc imgurService.AlbumService, repo repository.SkateparkRepository) SkateparksService {
	return &SkateparksServiceImpl{
		repo:   repo,
		imgSvc: imgSvc,
	}
}

// GetSkateparks gets the full list of skateparks from json repository
func (svc *SkateparksServiceImpl) GetSkateparks() (domain.Skateparks, error) {
	if svc.parkscache.skateparks == nil || len(svc.parkscache.skateparks) == 0 || time.Now().UTC().Sub(svc.parkscache.refreshedat).Hours() > 24 {
		fp := os.Getenv("SKATEPARKS_FILE")
		m, err := svc.repo.ParseSkateparks(fp)
		if err != nil {
			return nil, err
		}

		for i := range m {
			m[i].Album, err = svc.imgSvc.GetAlbum(m[i].AlbumID)
			if err != nil {
				log.Errorf("Error getting album:\n %v", err)
			}
		}

		svc.parkscache.skateparks = domain.Skateparks(m)
		svc.parkscache.refreshedat = time.Now().UTC()
	}

	return svc.parkscache.skateparks, nil
}
