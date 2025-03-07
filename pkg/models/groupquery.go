package models

import (
	"github.com/google/uuid"
)

type GroupQuery struct {
	ID               uuid.UUID          `json:"id" gorm:"type:uuid;primaryKey"`
	Name             string             `json:"name"`
	TagsInGroupQuery []TagsInGroupQuery `json:"tagsInGroupQueries" gorm:"foreignKey:GroupQueryID"`
}
