package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID                          uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey"`
	DeviceID                    uuid.UUID         `json:"device_id"`
	Name                        string            `json:"name"`
	Comment                     string            `json:"comment"`
	Active                      bool              `json:"active"`
	Address                     int               `json:"address"`
	Type                        int               `json:"type"` // Enum
	Recalc                      bool              `json:"recalc"`
	Factor                      float32           `json:"factor"`
	Offset                      float32           `json:"offset"`
	UseDeviceOrderBytes         bool              `json:"use_device_order_bytes"`
	OrderBytesID                uuid.UUID         `json:"order_bytes_id"`
	Value                       string            `json:"value"`
	HasWriteRegister            bool              `json:"has_write_register"`
	WriteRegisterAddress        int               `json:"write_register_address"`
	IsBit                       bool              `json:"is_bit"`
	BitIndex                    byte              `json:"bit_index"`
	ReadAfterWrite              bool              `json:"read_after_write"`
	AreaInfoID                  uuid.UUID         `json:"area_info_id"`
	Device                      Device            `json:"device" gorm:"foreignKey:DeviceID"`
	AreaInfo                    AreaInfo          `json:"area_info" gorm:"foreignKey:AreaInfoID"`
	OrderBytes                  OrderBytes        `json:"order_bytes" gorm:"foreignKey:OrderBytesID"`
	TagsInGroupQuery            *TagsInGroupQuery `json:"tags_in_group_query" gorm:"foreignKey:TagsInGroupQueryID"`
	TagsInGroupQueryID          uuid.UUID         `json:"tags_in_group_query_id"`
	OutputType                  int               `json:"output_type"` // Enum
	UseOutputType               bool              `json:"use_output_type"`
	RoundingAccuracy            byte              `json:"rounding_accuracy"`
	ConvertBeforeRecalcForRead  bool              `json:"convert_before_recalc_for_read"`
	ConvertBeforeRecalcForWrite bool              `json:"convert_before_recalc_for_write"`
}
