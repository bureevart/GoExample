package enums

type BlocksCreationAlgorithm int

const (
	MaxBlockSizeBased              BlocksCreationAlgorithm = iota // 0
	WithoutReadEmptyAddressesBased                                // 1
)
