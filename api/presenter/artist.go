package presenter

import "go-dummy-clean-arch/entity"

type Artist struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (a *Artist) FromEntity(artist *entity.Artist) {
	a.ID = artist.ID
	a.Name = artist.Name
}

func (a Artist) ToEntity() *entity.Artist {
	return &entity.Artist{
		ID:   a.ID,
		Name: a.Name,
	}
}
