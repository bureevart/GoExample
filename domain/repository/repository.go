package repository

import (
	"context"

	"goexample/domain/models"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetAll(ctx context.Context) ([]TEntity, error)
}

type AreaInfoRepository interface {
	BaseRepository[models.AreaInfo]
}
