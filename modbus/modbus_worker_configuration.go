package modbus

import (
	"goexample/pkg/enums"

	"github.com/google/uuid"
)

// Node представляет узел (Node) в системе
type Node struct {
	DeviceId          uuid.UUID
	Port              int32
	Server            string
	UnitIdentifier    byte
	DeviceStartsFrom  int32
	DelayBetweenPolls int
	ReconnectTimeout  int
	ReadTimeout       int
	WriteTimeout      int
	ReadType          enums.ReadType
}
