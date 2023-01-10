package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrWrongCoinNumber = sdkerrors.Register(ModuleName, 2, "wrong number of coins")
var ErrInsufficientLiquidity = sdkerrors.Register(ModuleName, 3, "insufficient initial liquidity")
var ErrSmallOutCoin = sdkerrors.Register(ModuleName, 4, "calculated out coin is smaller than the minimum")
