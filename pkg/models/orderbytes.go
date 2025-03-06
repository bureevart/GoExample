package models

import (
	"github.com/google/uuid"
)

type OrderBytes struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name    string    `json:"name"`
	Comment string    `json:"comment"`
	Uint16  string    `json:"uint16"`
	Uint32  string    `json:"uint32"`
	Float   string    `json:"float"`
	Double  string    `json:"double"`
	String  string    `json:"string"`
	Code    int64     `json:"code"`
	Tags    []Tag     `json:"tags" gorm:"foreignKey:OrderBytesID"`
	Devices []Device  `json:"devices" gorm:"foreignKey:OrderBytesID"`
}
