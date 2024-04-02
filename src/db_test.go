package skatepark_api_test

import (
	"testing"

	skatepark_api "github.com/derekpedersen/skatepark-api-go/src"
)

const testfile = "../.tools/test-data/good/"

func TestSkateparks_GetSkateparks_Success(t *testing.T) {
	// Arrange

	// Act
	parks, err := skatepark_api.ParseSkateparks(testfile)

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
	fp := "../.tools/test-data/bad/"

	// Act
	_, err := skatepark_api.ParseSkateparks(fp)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestSkateparks_GetSkateparks_Error_Parse(t *testing.T) {
	// Arrange
	fp := "../.tools/test-data/bad/"

	// Act
	_, err := skatepark_api.ParseSkateparks(fp)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}
