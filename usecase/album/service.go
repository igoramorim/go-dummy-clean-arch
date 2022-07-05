package album

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

func (s *Service) GetByID(ctx context.Context, id int64) (entity.Album, error) {
	log.Println("Album Service : GetByID:", id)
	return s.repo.GetByID(ctx, id)
}
