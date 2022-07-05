package album

import (
	"context"
	"go-dummy-clean-arch/entity"
)

type UseCase interface {
	GetByID(ctx context.Context, id int64) (entity.Album, error)
}

type Repository interface {
	GetByID(ctx context.Context, id int64) (entity.Album, error)
}
