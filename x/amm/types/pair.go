package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"strconv"
)

const (
	ReserveAddressPrefix = "reserve/"
	ShareDenomPrefix     = "amm/share/"
)

func NewPair(pairID uint64, denom0, denom1 string) Pair {
	return Pair{
		Id:     pairID,
		Denom0: denom0,
		Denom1: denom1,
	}
}

func SortDenoms(denomA, denomB string) (denom0, denom1 string) {
	if denomA < denomB {
		return denomA, denomB
	}

	return denomB, denomA
}

func PairReserveAddress(pair Pair) sdk.AccAddress {
	return address.Module(
		ModuleName, []byte(ReserveAddressPrefix+strconv.FormatUint(pair.Id, 10)))
}

func ShareDenom(pair Pair) string {
	return ShareDenomPrefix + strconv.FormatUint(pair.Id, 10)
}
