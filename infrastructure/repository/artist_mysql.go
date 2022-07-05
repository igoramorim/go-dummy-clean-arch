package repository

import (
	"context"
	"database/sql"
	"go-dummy-clean-arch/entity"
	"log"
)

type artistMySQL struct {
	db *sql.DB
}

func NewArtistMysql(db *sql.DB) *artistMySQL {
	return &artistMySQL{
		db: db,
	}
}

func (r *artistMySQL) GetByID(ctx context.Context, id int64) (entity.Artist, error) {
	log.Println("artistMySQL : GetByID :", id)

	var artist entity.Artist

	row := r.db.QueryRow("select * from artist where id = ?", id)
	if err := row.Scan(&artist.ID, &artist.Name); err != nil {
		if err == sql.ErrNoRows {
			return entity.Artist{}, err
		}
		return entity.Artist{}, err
	}

	return artist, nil
}

func (r *artistMySQL) Save(ctx context.Context, artist entity.Artist) error {
	log.Println("artistMySQL : Save :", artist)

	result, err := r.db.Exec("insert into artist (name) values (?)", artist.Name)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
