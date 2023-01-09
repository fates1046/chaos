package types

const (
	// ModuleName defines the module name
	ModuleName = "amm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_amm"
)

var (
	LastPairIDKey      = []byte{0x01}
	PairKeyPrefix      = []byte{0x02}
	PairIndexKeyPrefix = []byte{0x03}
)
