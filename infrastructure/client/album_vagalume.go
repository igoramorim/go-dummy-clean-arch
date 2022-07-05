package client

import (
	"context"
	"go-dummy-clean-arch/entity"
	"log"
)

type vagalumeAlbumClient struct {
	// some fields
}

func NewVagalumeAlbumCLient() *vagalumeAlbumClient {
	return &vagalumeAlbumClient{}
}

func (c *vagalumeAlbumClient) GetByID(ctx context.Context, id int64) (entity.Album, error) {
	log.Println("vagalumeAlbumClient : GetByID :", id)
	return entity.Album{}, nil
}
