package repository

import (
	"testing"
)

func TestSkateparks_GetSkateparks_Success(t *testing.T) {
	// Arrange
	fp := "../testutils/skateparks.json"

	repo := NewSkateparkRepository()

	// Act
	parks, err := repo.ParseSkateparks(fp)

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
	fp := "../testutils/iamnotafile.json"

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
	fp := "../testutils/bad.json"

	repo := NewSkateparkRepository()

	// Act
	_, err := repo.ParseSkateparks(fp)

	// Assert
	if err == nil {
		t.Fatal("expected an error")
	}
}
