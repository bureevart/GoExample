package models

import (
	"github.com/google/uuid"
)

// Node представляет узел (Node) в системе
type Node struct {
	ID                  uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name                string    `json:"name"`
	Comment             string    `json:"comment"`
	Active              bool      `json:"active"`
	Address             string    `json:"address"`
	Port                string    `json:"port"`
	DisconnectEveryPoll bool      `json:"disconnectEveryPoll"`
	AutoStart           bool      `json:"autoStart"`
	Devices             []Device  `json:"devices" gorm:"foreignKey:NodeID"`
}
