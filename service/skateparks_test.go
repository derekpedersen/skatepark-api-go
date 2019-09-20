package service_test

import (
	"fmt"
	"os"
	"testing"

	imgModel "github.com/derekpedersen/imgur-go/model"
	"github.com/derekpedersen/skatepark-api-go/mock"
	"github.com/derekpedersen/skatepark-api-go/repository"
	"github.com/derekpedersen/skatepark-api-go/service"
	"github.com/golang/mock/gomock"
)

func TestSkateparksService_GetSkateparks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService(ctrl)
	imgSvc.EXPECT().GetAlbum(gomock.Any()).Return(&imgModel.Album{}, nil).Times(2)
	os.Setenv("SKATEPARKS_FILE", "../testutils/skateparks.json")

	repo := repository.NewSkateparkRepository()
	svc := service.NewSkateparksService(imgSvc, repo)

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

func TestSkateparks_GetSkateparks_Error_FileRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService(ctrl)
	os.Setenv("SKATEPARKS_FILE", "../testutils/iamnotafile.json")

	repo := repository.NewSkateparkRepository()
	svc := service.NewSkateparksService(imgSvc, repo)

	// Act
	_, err := svc.GetSkateparks()

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestSkateparks_GetSkateparks_Error_Parse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService(ctrl)
	os.Setenv("SKATEPARKS_FILE", "../testutils/bad.json")

	repo := repository.NewSkateparkRepository()
	svc := service.NewSkateparksService(imgSvc, repo)

	// Act
	_, err := svc.GetSkateparks()

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestSkateparksService_GetSkateparks_Error_Imgur(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService(ctrl)
	imgSvc.EXPECT().GetAlbum(gomock.Any()).Return(nil, fmt.Errorf("imgur is down")).Times(2)
	os.Setenv("SKATEPARKS_FILE", "../testutils/skateparks.json")

	repo := repository.NewSkateparkRepository()
	svc := service.NewSkateparksService(imgSvc, repo)

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
