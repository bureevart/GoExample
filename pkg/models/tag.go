package models

import (
	"goexample/pkg/enums"

	"github.com/google/uuid"
)

type Tag struct {
	ID                          uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey"`
	DeviceID                    uuid.UUID         `json:"deviceId"`
	Name                        string            `json:"name"`
	Comment                     string            `json:"comment"`
	Active                      bool              `json:"active"`
	Address                     int               `json:"address"`
	Type                        enums.Type        `json:"type"` // Enum
	Recalc                      bool              `json:"recalc"`
	Factor                      float32           `json:"factor"`
	Offset                      float32           `json:"offset"`
	UseDeviceOrderBytes         bool              `json:"useDeviceOrderBytes"`
	OrderBytesID                uuid.UUID         `json:"orderBytesId"`
	Value                       string            `json:"value"`
	HasWriteRegister            bool              `json:"hasWriteRegister"`
	WriteRegisterAddress        int               `json:"writeRegisterAddress"`
	IsBit                       bool              `json:"isBit"`
	BitIndex                    byte              `json:"bitIndex"`
	ReadAfterWrite              bool              `json:"readAfterWrite"`
	AreaInfoID                  uuid.UUID         `json:"areaInfoId"`
	Device                      *Device           `json:"device" gorm:"foreignKey:DeviceID"`
	AreaInfo                    *AreaInfo         `json:"areaInfo" gorm:"foreignKey:AreaInfoID"`
	OrderBytes                  *OrderBytes       `json:"orderBytes" gorm:"foreignKey:OrderBytesID"`
	TagsInGroupQuery            *TagsInGroupQuery `json:"tagsInGroupQuery" gorm:"foreignKey:TagsInGroupQueryID"`
	TagsInGroupQueryID          uuid.UUID         `json:"tagsInGroupQueryId"`
	OutputType                  int               `json:"outputType"` // Enum
	UseOutputType               bool              `json:"useOutputType"`
	RoundingAccuracy            byte              `json:"roundingAccuracy"`
	ConvertBeforeRecalcForRead  bool              `json:"convertBeforeRecalcForRead"`
	ConvertBeforeRecalcForWrite bool              `json:"convertBeforeRecalcForWrite"`
}
