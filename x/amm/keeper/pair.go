package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fates1046/chaos/x/amm/types"
)

func (k Keeper) AddLiquidity(ctx sdk.Context, fromAddr sdk.AccAddress, coins sdk.Coins) (mintedShare sdk.Coin, err error) {
	coin0, coin1 := coins[0], coins[1]

	pair, found := k.GetPairByDenoms(ctx, coin0.Denom, coin1.Denom)

	// empty liquidity pool
	if !found {
		pairID := k.GetLastPairID(ctx)
		pairID++
		k.SetLastPairID(ctx, pairID)

		pair = types.NewPair(pairID, coin0.Denom, coin1.Denom)
		k.SetPair(ctx, pair)
		k.SetPairIndex(ctx, pair)
	}

	reserveAddr := types.PairReserveAddress(pair)
	shareDenom := types.ShareDenom(pair)

	reserveBalances := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
	rx := reserveBalances.AmountOf(coin0.Denom)
	ry := reserveBalances.AmountOf(coin1.Denom)
	x := coin0.Amount
	y := coin1.Amount
	totalShare := k.bankKeeper.GetSupply(ctx, shareDenom).Amount

	var ax, ay, share sdk.Int
	if totalShare.IsZero() {
		var l sdk.Dec
		l, err = sdk.NewDecFromInt(x.Mul(y)).ApproxSqrt()
		if err != nil {
			return
		}
		share = l.TruncateInt()
		if share.LT(k.GetMinInitialLiquidity(ctx)) {
			err = sdkerrors.Wrapf(types.ErrInsufficientLiquidity, "insufficient initial liquidity: %s", share)
			return
		}

		ax = x
		ay = y
	} else {
		share = sdk.MinInt(x.Quo(rx), y.Quo(ry)).Mul(totalShare)
		ax = rx.Mul(share).Quo(totalShare)
		ay = ry.Mul(share).Quo(totalShare)
	}

	if !ax.IsPositive() || !ay.IsPositive() || !share.IsPositive() {
		err = types.ErrInsufficientLiquidity
		return
	}

	amt := sdk.NewCoins(sdk.NewCoin(coin0.Denom, ax), sdk.NewCoin(coin1.Denom, ay))
	if err = k.bankKeeper.SendCoins(ctx, fromAddr, reserveAddr, amt); err != nil {
		return
	}
	mintedShare = sdk.NewCoin(shareDenom, share)
	if err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mintedShare)); err != nil {
		return
	}
	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fromAddr, sdk.NewCoins(mintedShare)); err != nil {
		return
	}
	return mintedShare, nil

}

func (k Keeper) RemoveLiquidity(ctx sdk.Context, fromAddr sdk.AccAddress, share sdk.Coin) (withdrawnCoins sdk.Coins, err error) {
	var pairID uint64
	pairID, err = types.ParseShareDenom(share.Denom)
	if err != nil {
		return
	}

	pair, found := k.GetPair(ctx, pairID)
	if !found { // Sanity check
		err = sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pair %d not found", pairID)
		return
	}

	reserveAddr := types.PairReserveAddress(pair)
	reserveBalances := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
	rx := reserveBalances.AmountOf(pair.Denom0)
	ry := reserveBalances.AmountOf(pair.Denom1)
	totalShare := k.bankKeeper.GetSupply(ctx, share.Denom).Amount

	var wx, wy sdk.Int
	if share.Amount.Equal(totalShare) {
		wx = rx
		wy = ry
		k.DeletePair(ctx, pair)
		k.DeletePairIndex(ctx, pair)
	} else {
		wx = rx.Mul(share.Amount).Quo(totalShare)
		wy = ry.Mul(share.Amount).Quo(totalShare)
	}
	if !wx.IsPositive() && !wy.IsPositive() {
		err = sdkerrors.Wrap(types.ErrInsufficientLiquidity, "too small share to remove")
		return
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, fromAddr, types.ModuleName, sdk.NewCoins(share)); err != nil {
		return
	}

	if err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(share)); err != nil {
		return
	}
	withdrawnCoins = sdk.NewCoins(sdk.NewCoin(pair.Denom0, wx), sdk.NewCoin(pair.Denom1, wy))
	if err = k.bankKeeper.SendCoins(ctx, reserveAddr, fromAddr, withdrawnCoins); err != nil {
		return
	}
	return withdrawnCoins, nil

}
