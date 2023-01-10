package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fates1046/chaos/x/amm/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetAddLiquidityCmd(),
		GetRemovedLiquidityCmd(),
	)

	return cmd
}

func GetAddLiquidityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-liquidity [coins]",
		Args:  cobra.ExactArgs(1),
		Short: "Add liquidity to a pair",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return fmt.Errorf("invalid coins: %w", err)
			}
			msg := types.NewMsgAddLiquidity(clientCtx.GetFromAddress(), coins)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetRemovedLiquidityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-liquidity [share]",
		Args:  cobra.ExactArgs(1),
		Short: "Remove liquidity from a pair",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			share, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return fmt.Errorf("invalid share: %w", err)
			}
			msg := types.NewMsgRemoveLiquidity(clientCtx.GetFromAddress(), share)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
