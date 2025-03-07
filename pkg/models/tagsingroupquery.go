package models

import (
	"github.com/google/uuid"
)

type TagsInGroupQuery struct {
	ID           uuid.UUID   `json:"id" gorm:"type:uuid;primaryKey"`
	TagID        uuid.UUID   `json:"tagId"`
	GroupQueryID uuid.UUID   `json:"groupQueryId"`
	GroupQuery   *GroupQuery `json:"groupQuery" gorm:"foreignKey:GroupQueryID"`
	Tag          *Tag        `json:"tag" gorm:"foreignKey:TagID"`
}
