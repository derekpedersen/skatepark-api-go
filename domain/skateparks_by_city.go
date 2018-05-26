package domain

import "github.com/derekpedersen/skatepark-api-go/model"

type SkateparksByCity []model.SkateparkByCity

func (s SkateparksByCity) Len() int {
	return len(s)
}

func (s SkateparksByCity) Less(i, j int) bool {
	return s[i].City < s[j].City
}

func (s SkateparksByCity) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
