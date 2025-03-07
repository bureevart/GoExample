package usecase

import (
	"context"

	"goexample/config"
	"goexample/domain/models"
	"goexample/domain/repository"
	"goexample/usecase/dto"
)

type AreaInfoUsecase struct {
	base *BaseUsecase[models.AreaInfo, dto.CreateAreainfo, dto.UpdateAreainfo, dto.AreaInfo]
}

func NewAreaInfoUsecase(cfg *config.Config, repository repository.AreaInfoRepository) *AreaInfoUsecase {
	return &AreaInfoUsecase{
		base: NewBaseUsecase[models.AreaInfo, dto.CreateAreainfo, dto.UpdateAreainfo, dto.AreaInfo](cfg, repository),
	}
}

// Create
func (u *AreaInfoUsecase) Create(ctx context.Context, req dto.CreateAreainfo) (dto.AreaInfo, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *AreaInfoUsecase) Update(ctx context.Context, id int, req dto.UpdateAreainfo) (dto.AreaInfo, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *AreaInfoUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *AreaInfoUsecase) GetById(ctx context.Context, id int) (dto.AreaInfo, error) {
	return s.base.GetById(ctx, id)
}

// Get By Id
func (s *AreaInfoUsecase) GetAll(ctx context.Context, id int) (dto.AreaInfo, error) {
	return s.base.GetAll(ctx)
}
