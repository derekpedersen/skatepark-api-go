package domain

import "github.com/derekpedersen/skatepark-api-go/model"

// SkateparksByCity is a collection of skateparks grouped by city
type SkateparksByCity []model.SkateparkByCity

func (sk SkateparksByCity) Len() int           { return len(sk) }
func (sk SkateparksByCity) Swap(i, j int)      { sk[i], sk[j] = sk[j], sk[i] }
func (sk SkateparksByCity) Less(i, j int) bool { return sk[i].City < sk[j].City }
