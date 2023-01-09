package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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

func GetPairKey(pairID uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(pairID)...)
}

func GetPairIndexKey(denom0, denom1 string) []byte {
	return append(append(PairIndexKeyPrefix, LengthPrefix([]byte(denom0))...), denom1...)
}

func LengthPrefix(b []byte) []byte {
	l := len(b)
	if l == 0 {
		return b
	}
	return append([]byte{byte(l)}, b...)
}
