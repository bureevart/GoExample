package models

import (
	"github.com/google/uuid"
)

type AreaInfo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        int64     `json:"code"`
	Tags        []Tag     `json:"tags" gorm:"foreignKey:AreaInfoID"`
}
