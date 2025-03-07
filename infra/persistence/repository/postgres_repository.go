package repository

import (
	"context"

	"goexample/common"
	"goexample/config"
	database "goexample/infra/persistence/database"
	"goexample/pkg/logging"

	"gorm.io/gorm"
)

const softDeleteExp string = "id = ? and deleted_by is null"

type BaseRepository[TEntity any] struct {
	database *gorm.DB
	logger   logging.Logger
	preloads []database.PreloadEntity
}

func NewBaseRepository[TEntity any](cfg *config.Config, preloads []database.PreloadEntity) *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		database: database.GetDb(),
		logger:   logging.NewLogger(cfg),
		preloads: preloads,
	}
}

func (r BaseRepository[TEntity]) Create(ctx context.Context, entity TEntity) (TEntity, error) {
	tx := r.database.WithContext(ctx).Begin()
	err := tx.
		Create(&entity).
		Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return entity, err
	}
	tx.Commit()

	return entity, nil
}

func (r BaseRepository[TEntity]) Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error) {
	snakeMap := map[string]interface{}{}
	for k, v := range entity {
		snakeMap[common.ToSnakeCase(k)] = v
	}

	model := new(TEntity)
	tx := r.database.WithContext(ctx).Begin()

	if err := tx.Model(model).
		Where("id = ?", id).
		Updates(snakeMap).
		Error; err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return *model, err
	}

	tx.Commit()
	return *model, nil
}

func (r BaseRepository[TEntity]) Delete(ctx context.Context, id int) error {
	tx := r.database.WithContext(ctx).Begin()

	model := new(TEntity)

	// Удаляем запись из БД полностью
	if err := tx.Where("id = ?", id).Delete(model).Error; err != nil {
		tx.Rollback()
		r.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err
	}

	tx.Commit()
	return nil
}

func (r BaseRepository[TEntity]) GetById(ctx context.Context, id int) (TEntity, error) {
	model := new(TEntity)
	db := database.Preload(r.database, r.preloads)
	err := db.
		Where(softDeleteExp, id).
		First(model).
		Error
	if err != nil {
		return *model, err
	}
	return *model, nil
}

func (r BaseRepository[TEntity]) GetAll(ctx context.Context) ([]TEntity, error) {
	var models []TEntity
	db := database.Preload(r.database, r.preloads)
	err := db.
		Where("deleted_by IS NULL"). // Только активные записи
		Find(&models).
		Error
	if err != nil {
		return nil, err
	}
	return models, nil
}
