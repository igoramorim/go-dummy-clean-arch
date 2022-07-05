package artist

import (
	"context"
	"go-dummy-clean-arch/entity"
)

type UseCase interface {
	GetByID(ctx context.Context, id int64) (entity.Artist, error)
	Save(ctx context.Context, artist entity.Artist) error
}

type Repository interface {
	GetByID(ctx context.Context, id int64) (entity.Artist, error)
	Save(ctx context.Context, artist entity.Artist) error
}
