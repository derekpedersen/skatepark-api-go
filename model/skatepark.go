package model

import imgurModel "github.com/derekpedersen/imgur-go/model"

// Skatepark model
// swagger:model
type Skatepark struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Address     Address           `json:"address"`
	AlbumID     string            `json:"albumId"`
	Album       *imgurModel.Album `json:"album"`
}
