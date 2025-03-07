package models

import (
	"github.com/google/uuid"
)

type GroupQuery struct {
	ID               uuid.UUID          `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string             `json:"name"`
	TagsInGroupQuery []TagsInGroupQuery `json:"tags_in_group_queries" gorm:"foreignKey:GroupQueryID"`
}
