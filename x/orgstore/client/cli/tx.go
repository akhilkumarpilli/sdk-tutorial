package cli

import (
	"fmt"
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	orgstoreTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Orgstore transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	orgstoreTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateOrg(cdc),
		GetCmdDeleteOrg(cdc),
		GetCmdAddUser(cdc),
		GetCmdDeleteUser(cdc),
	)...)

	return orgstoreTxCmd
}

// GetCmdCreateOrg is the CLI command for sending a CreateOrg transaction
func GetCmdCreateOrg(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-org [org-name]",
		Short: "create new organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCreateOrg(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteOrg is the CLI command for deleting organization
func GetCmdDeleteOrg(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-org [org-name]",
		Short: "delete the organization that you own",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			msg := types.NewMsgDeleteOrg(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddUser is the CLI command for sending a AddUser transaction
func GetCmdAddUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-user [org-name] [user-name] [user-addr] [user-role]",
		Short: "add the user to the org that you own",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			userAddr, convertErr := sdk.AccAddressFromBech32(args[2])
			if convertErr != nil {
				fmt.Print(convertErr)
				return convertErr
			}

			msg := types.NewMsgAddUser(args[0], cliCtx.GetFromAddress(), args[1], userAddr, args[3])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteUser is the CLI command for sending a DeleteName transaction
func GetCmdDeleteUser(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-user [org-name] [user-addr]",
		Short: "delete the user from the org that you own along with it's associated fields",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			userAddr, convertErr := sdk.AccAddressFromBech32(args[1])
			if convertErr != nil {
				fmt.Print(convertErr)
				return convertErr
			}

			msg := types.NewMsgDeleteUser(args[0], cliCtx.GetFromAddress(), userAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
