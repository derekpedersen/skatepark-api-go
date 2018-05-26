package domain

import "github.com/derekpedersen/skatepark-api-go/model"

type SkateparksByState []model.SkateparkByState

func (s SkateparksByState) Len() int {
	return len(s)
}

func (s SkateparksByState) Less(i, j int) bool {
	return s[i].State < s[j].State
}

func (s SkateparksByState) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
