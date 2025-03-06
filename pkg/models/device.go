package models

import (
	"goexample/pkg/enums" // Путь до вашего пакета

	"github.com/google/uuid"
)

type Device struct {
	ID                      uuid.UUID                     `json:"id" gorm:"type:uuid;primaryKey"`
	NodeID                  uuid.UUID                     `json:"node_id"`
	Name                    string                        `json:"name"`
	Comment                 string                        `json:"comment"`
	Active                  bool                          `json:"active"`
	Address                 byte                          `json:"address"`
	ResponseTimeout         int                           `json:"response_timeout"`
	ReconnectTimeout        int                           `json:"reconnect_timeout"`
	RetryError              int                           `json:"retry_error"`
	RetryWriteError         int                           `json:"retry_write_error"`
	PollPeriod              int                           `json:"poll_period"`
	AutoStart               bool                          `json:"auto_start"`
	MaxRegistersRequest     byte                          `json:"max_registers_request"`
	OrderBytesID            uuid.UUID                     `json:"order_bytes_id"`
	Tags                    []Tag                         `json:"tags" gorm:"foreignKey:DeviceID"`
	Node                    Node                          `json:"node" gorm:"foreignKey:NodeID"`
	OrderBytes              OrderBytes                    `json:"order_bytes" gorm:"foreignKey:OrderBytesID"`
	PollDeactivatedTags     bool                          `json:"poll_deactivated_tags"`
	BlocksCreationAlgorithm enums.BlocksCreationAlgorithm `json:"blocks_creation_algorithm"` // Enum
}
