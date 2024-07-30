package skatepark_api_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/derekpedersen/imgur-go/album"
	"github.com/derekpedersen/imgur-go/authorization"
	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
)

const testfile = "../.db/test/good/"

func TestSkateparks_GetSkateparks_Success(t *testing.T) {

	t.Skip("need to mock out imgur album call")

	// Arrange

	imgurAuth, err := authorization.NewAuthorization()
	if err != nil {
		log.Fatal(err)
	}

	if skatepark_api.IMGUR_SVC = album.NewAlbumService(*imgurAuth, "https://api.imgur.com/3/album/"); skatepark_api.IMGUR_SVC == nil {
		log.Fatal(fmt.Errorf("failed to load imgur service"))
	}

	// Act
	parks, err := skatepark_api.ParseSkateparks(testfile, true)

	// Assert
	if err != nil {
		t.Fatalf("encountered an unexpected error: %v", err)
	}
	if len(parks) == 0 {
		t.Errorf("expected a result length greater than 0")
	}
}

func TestSkateparks_GetSkateparks_Error_FileRead(t *testing.T) {
	// Arrange
	fp := "../.db/test/bad/"

	// Act
	_, err := skatepark_api.ParseSkateparks(fp, false)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestSkateparks_GetSkateparks_Error_Parse(t *testing.T) {
	// Arrange
	fp := "../.dev/test/bad/"

	// Act
	_, err := skatepark_api.ParseSkateparks(fp, false)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}
