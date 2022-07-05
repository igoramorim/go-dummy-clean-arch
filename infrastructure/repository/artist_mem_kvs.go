package repository

import (
	"context"
	"errors"
	"fmt"
	"go-dummy-clean-arch/entity"
	"log"
)

type artistMemoryKVS struct {
	artists map[int64]entity.Artist
}

func NewArtistMemoryKVS() *artistMemoryKVS {
	return &artistMemoryKVS{
		artists: make(map[int64]entity.Artist),
	}
}

func (r *artistMemoryKVS) GetByID(ctx context.Context, id int64) (entity.Artist, error) {
	log.Println("artistMemoryKVS : GetByID :", id)
	return r.artists[id], nil
}

func (r *artistMemoryKVS) Save(ctx context.Context, artist entity.Artist) error {
	log.Println("artistMemoryKVS : Save :", artist)
	_, exists := r.artists[artist.ID]
	if exists {
		return errors.New(fmt.Sprintf("artist %v already exists", artist))
	}

	r.artists[artist.ID] = artist
	return nil
}
