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
	tests := []struct {
		Name string
	}{
		{
			Name: "Constructor",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	imgSvc := mock.NewMockAlbumService()
	imgSvc.EXPECT().GetAlbum(gomock.Any()).Return(imgModel.Album{})
	os.Setenv("SKATEPARKS_FILE", "../repository/json/skateparks.json")

	repo := repository.NewSkateparkRepository()
	svc := NewSkateparksService(imgSvc, repo)

	for _, test := range tests {
		// Act
		parks, err := svc.GetSkateparks()

		// Assert
		if err != nil {
			t.Fatalf("%v, encountered an unexpected error: %v", test.Name, err)
		}
		if len(parks) == 0 {
			t.Errorf("%v: expected a result length greater than 0", test.Name)
		}
	}
}
