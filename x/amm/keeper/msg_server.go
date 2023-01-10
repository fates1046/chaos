package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fates1046/chaos/x/amm/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func (m msgServer) AddLiquidity(c context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	mintedShare, err := m.Keeper.AddLiquidity(
		ctx, sdk.MustAccAddressFromBech32(msg.Sender), msg.Coins,
	)
	if err != nil {
		return nil, err
	}
	return &types.MsgAddLiquidityResponse{
		MintedShare: mintedShare,
	}, nil

}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}
