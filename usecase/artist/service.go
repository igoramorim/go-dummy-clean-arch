package artist

import (
	"context"
	"go-dummy-clean-arch/entity"
	"log"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByID(ctx context.Context, id int64) (entity.Artist, error) {
	log.Println("Artist Service : GetByID:", id)
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Save(ctx context.Context, artist entity.Artist) error {
	log.Println("Artist Service : Save:", artist)
	return s.repo.Save(ctx, artist)
}
