package repository

import (
	"testing"
)

const testfile = "../.tools/test-data/good/"

func TestSkateparks_GetSkateparks_Success(t *testing.T) {
	// Arrange
	repo := NewSkateparkRepository()

	// Act
	parks, err := repo.ParseSkateparks(testfile)

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
	fp := "../.tools/"

	repo := NewSkateparkRepository()

	// Act
	_, err := repo.ParseSkateparks(fp)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestSkateparks_GetSkateparks_Error_Parse(t *testing.T) {
	// Arrange
	fp := "../.tools/test-data/bad/"

	repo := NewSkateparkRepository()

	// Act
	_, err := repo.ParseSkateparks(fp)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}
