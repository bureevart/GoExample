package enums

type ReadType int

const (
	NoneReadType ReadType = iota
	Coils
	DiscreteInputs
	HoldingRegisters
	InputRegisters
)
