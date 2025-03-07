package models

import (
	"goexample/pkg/enums" // Путь до вашего пакета

	"github.com/google/uuid"
)

type Device struct {
	ID                      uuid.UUID                     `json:"id" gorm:"type:uuid;primaryKey"`
	NodeID                  uuid.UUID                     `json:"nodeId"`
	Name                    string                        `json:"name"`
	Comment                 string                        `json:"comment"`
	Active                  bool                          `json:"active"`
	Address                 byte                          `json:"address"`
	ResponseTimeout         int                           `json:"responseTimeout"`
	ReconnectTimeout        int                           `json:"reconnectTimeout"`
	RetryError              int                           `json:"retryError"`
	RetryWriteError         int                           `json:"retryWriteError"`
	PollPeriod              int                           `json:"pollPeriod"`
	AutoStart               bool                          `json:"autoStart"`
	MaxRegistersRequest     byte                          `json:"maxRegistersRequest"`
	OrderBytesID            uuid.UUID                     `json:"orderBytesId"`
	Tags                    []Tag                         `json:"tags" gorm:"foreignKey:DeviceID"`
	Node                    *Node                         `json:"node" gorm:"foreignKey:NodeID"`
	OrderBytes              *OrderBytes                   `json:"orderBytes" gorm:"foreignKey:OrderBytesID"`
	PollDeactivatedTags     bool                          `json:"pollDeactivatedTags"`
	BlocksCreationAlgorithm enums.BlocksCreationAlgorithm `json:"blocksCreationAlgorithm"` // Enum
}
