package repository

import (
	"context"
	"database/sql"
	"go-dummy-clean-arch/entity"
	"log"
)

type albumMySQL struct {
	db *sql.DB
}

func NewAlbumMysql(db *sql.DB) *albumMySQL {
	return &albumMySQL{
		db: db,
	}
}

func (r *albumMySQL) GetByID(ctx context.Context, id int64) (entity.Album, error) {
	log.Println("albumMySQL : GetByID :", id)
	return entity.Album{}, nil
}
