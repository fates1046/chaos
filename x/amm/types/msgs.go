package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAddLiquidity    = "add_liquidity"
	TypeMsgRemoveLiquidity = "remove_liquidity"
)

var (
	_ sdk.Msg = (*MsgAddLiquidity)(nil)
	_ sdk.Msg = (*MsgRemoveLiquidity)(nil)
)

func NewMsgAddLiquidity(sender sdk.AccAddress, coins sdk.Coins) *MsgAddLiquidity {
	return &MsgAddLiquidity{
		Sender: sender.String(),
		Coins:  coins,
	}
}

func NewMsgRemoveLiquidity(sender sdk.AccAddress, share sdk.Coin) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Sender: sender.String(),
		Share:  share,
	}
}

func (msg *MsgAddLiquidity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg *MsgAddLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address %s", err)
	}
	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, err.Error())
	}
	if len(msg.Coins) != 2 {
		return ErrWrongCoinNumber
	}

	return nil
}

func (msg *MsgAddLiquidity) Route() string { return RouterKey }
func (msg *MsgAddLiquidity) Type() string  { return TypeMsgAddLiquidity }

func (msg *MsgAddLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg *MsgRemoveLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if _, err := ParseShareDenom(msg.Share.Denom); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return nil
}

func (msg *MsgRemoveLiquidity) Route() string { return RouterKey }
func (msg *MsgRemoveLiquidity) Type() string  { return TypeMsgRemoveLiquidity }

func (msg *MsgRemoveLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
