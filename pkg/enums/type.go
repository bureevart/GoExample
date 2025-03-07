package enums

type Type int

const (
	NoneType Type = iota // 0
	Bool
	Int16
	uint16
	Int32
	Uint32
	Float
	Double
	String
	DateTime = 9
)
