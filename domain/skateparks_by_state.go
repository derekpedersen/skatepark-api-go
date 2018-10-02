package domain

import "github.com/derekpedersen/skatepark-api-go/model"

// SkateparksByState are a group of skateparks groupe by state then city
type SkateparksByState []model.SkateparkByState

func (sk SkateparksByState) Len() int           { return len(sk) }
func (sk SkateparksByState) Swap(i, j int)      { sk[i], sk[j] = sk[j], sk[i] }
func (sk SkateparksByState) Less(i, j int) bool { return sk[i].State < sk[j].State }
