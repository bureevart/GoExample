package models

import (
	"github.com/google/uuid"
)

type TagsInGroupQuery struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	TagID        uuid.UUID  `json:"tag_id"`
	GroupQueryID uuid.UUID  `json:"group_query_id"`
	GroupQuery   GroupQuery `json:"group_query" gorm:"foreignKey:GroupQueryID"`
	Tag          *Tag       `json:"tag" gorm:"foreignKey:TagID"`
}
