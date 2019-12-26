package keeper

import (
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryOrgData = "org_data"
	QueryOrgs    = "org_list"
	QueryUsers   = "org_users"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOrgData:
			return queryOrgData(ctx, path[1:], req, keeper)
		case QueryUsers:
			return queryUsers(ctx, path[1:], req, keeper)
		case QueryOrgs:
			return queryOrgs(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

// nolint: unparam
func queryUsers(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	data := keeper.GetOrgData(ctx, path[0])

	if data.Name == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, data.Users)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// nolint: unparam
func queryOrgData(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	data := keeper.GetOrgData(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, data)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryOrgs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var namesList types.QueryResOrgs

	iterator := keeper.GetOrgsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
