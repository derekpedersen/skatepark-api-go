package model

import imgurModel "github.com/derekpedersen/imgur-go/model"

// skatepark
type Skatepark struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     struct {
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		City         string `json:"city"`
		State        string `json:"state"`
		ZipCode      string `json:"zipCode"`
		GoogleMaps   struct {
			EmbedMap  string `json:"embedMap"`
			ShareLink string `json:"ShareLink"`
		} `json:"googleMaps"`
	} `json:"address"`
	AlbumID string            `json:"albumId"`
	Album   *imgurModel.Album `json:"album"`
}
