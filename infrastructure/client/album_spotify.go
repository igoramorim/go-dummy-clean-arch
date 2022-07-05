package client

import (
	"context"
	"go-dummy-clean-arch/entity"
	"log"
)

type spotifyAlbumClient struct {
	// some fields
}

func NewSpotifyAlbumClient() *spotifyAlbumClient {
	return &spotifyAlbumClient{}
}

func (c *spotifyAlbumClient) GetByID(ctx context.Context, id int64) (entity.Album, error) {
	log.Println("spotifyAlbumClient : GetByID :", id)
	return entity.Album{}, nil
}
