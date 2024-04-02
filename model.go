package skatepark_api

import "github.com/derekpedersen/imgur-go/album"

// swagger:model googleMaps
type GoogleMaps struct {
	EmbedMap  string `json:"embedMap"`
	ShareLink string `json:"ShareLink"`
}

// swagger:model address
type Address struct {
	AddressLine1 string     `json:"addressLine1"`
	AddressLine2 string     `json:"addressLine2"`
	City         string     `json:"city"`
	State        string     `json:"state"`
	ZipCode      string     `json:"zipCode"`
	GoogleMaps   GoogleMaps `json:"googleMaps"`
}

// swagger:model skatepark
type Skatepark struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Address     Address      `json:"address"`
	AlbumID     string       `json:"albumId"`
	Album       *album.Album `json:"album"`
}

// swagger:model skateparks
type Skateparks []Skatepark

func (dom Skateparks) Len() int           { return len(dom) }
func (dom Skateparks) Swap(i, j int)      { dom[i], dom[j] = dom[j], dom[i] }
func (dom Skateparks) Less(i, j int) bool { return dom[i].Address.State < dom[j].Address.State }
