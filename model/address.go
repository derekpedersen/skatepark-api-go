package model

// Address model
// swagger:model address
type Address struct {
	AddressLine1 string     `json:"addressLine1"`
	AddressLine2 string     `json:"addressLine2"`
	City         string     `json:"city"`
	State        string     `json:"state"`
	ZipCode      string     `json:"zipCode"`
	GoogleMaps   GoogleMaps `json:"googleMaps"`
}
