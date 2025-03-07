package dto

import (
	"goexample/usecase/dto"

	"github.com/google/uuid"
)

type CreateAreaInfoRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        int64     `json:"code"`
}

type UpdateAreaInfoRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        int64     `json:"code"`
}
type AreaInfoResponse struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        int64     `json:"code"`
	//Tags        []TagResponse     `json:"tags"`
}

func ToAreaInfoResponse(from dto.AreaInfo) AreaInfoResponse {
	return AreaInfoResponse{
		ID:          from.ID,
		Name:        from.Name,
		Description: from.Description,
		Code:        from.Code,
	}
}

func ToCreateAreaInfo(from CreateAreaInfoRequest) dto.CreateAreainfo {
	return dto.CreateAreainfo{
		ID:          from.ID,
		Name:        from.Name,
		Description: from.Description,
		Code:        from.Code,
	}
}

func ToUpdateAreaInfo(from UpdateAreaInfoRequest) dto.UpdateAreainfo {
	return dto.UpdateAreainfo{
		ID:          from.ID,
		Name:        from.Name,
		Description: from.Description,
		Code:        from.Code,
	}
}
