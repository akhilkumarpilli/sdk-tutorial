package cli

import (
	"fmt"
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	orgstoreQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the orgstore module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	orgstoreQueryCmd.AddCommand(client.GetCommands(
		GetCmdUsers(storeKey, cdc),
		GetCmdOrgData(storeKey, cdc),
		GetCmdOrgs(storeKey, cdc),
	)...)
	return orgstoreQueryCmd
}

// GetCmdUsers queries information about users
func GetCmdUsers(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "org_users [name]",
		Short: "Users of org",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org_users/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", name)
				return nil
			}

			var out []types.OrgUser
			cdc.MustUnmarshalJSON(res, &out)
			p := types.Printer{Response: out}
			return cliCtx.PrintOutput(p)
		},
	}
}

// GetCmdOrgData queries information about org data
func GetCmdOrgData(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "org_data [name]",
		Short: "Query data of org",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org_data/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", name)
				return nil
			}

			var out types.Org
			cdc.MustUnmarshalJSON(res, &out)
			p := types.Printer{Response: out}
			return cliCtx.PrintOutput(p)
		},
	}
}

// GetCmdOrgs queries a list of all orgs
func GetCmdOrgs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "orgs_list",
		Short: "Org List",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res,err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org_list", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			var out types.QueryResOrgs
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
