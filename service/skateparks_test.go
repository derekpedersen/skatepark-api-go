package service

import (
	"os"
	"testing"

	imgModel "github.com/derekpedersen/imgur-go/model"
	"github.com/derekpedersen/skatepark-api-go/mock"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/golang/mock/gomock"
)

func TestSkateparks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService(ctrl)
	imgSvc.EXPECT().GetAlbum(gomock.Any()).Return(&imgModel.Album{}, nil).Times(2)
	os.Setenv("SKATEPARKS_FILE", "../testutils/skateparks.json")

	repo := repository.NewSkateparkRepository()
	svc := NewSkateparksService(imgSvc, repo)

	// Act
	parks, err := svc.GetSkateparks()

	// Assert
	if err != nil {
		t.Fatalf("encountered an unexpected error: %v", err)
	}
	if len(parks) == 0 {
		t.Errorf("expected a result length greater than 0")
	}
}
