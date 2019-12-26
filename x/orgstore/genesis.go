package orgstore

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	OrgRecords []Org `json:"whois_records"`
}

func NewGenesisState(orgRecords []Org) GenesisState {
	return GenesisState{OrgRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.OrgRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid OrgRecord: Value: %s. Error: Missing Owner", record.Name)
		}
		if record.Name == "" {
			return fmt.Errorf("invalid OrgRecord: Owner: %s. Error: Missing Name", record.Owner)
		}

	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		OrgRecords: []Org{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.OrgRecords {
		keeper.SetOrg(ctx, record.Name, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Org
	iterator := k.GetOrgsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		org := k.GetOrgData(ctx, name)
		records = append(records, org)

	}
	return GenesisState{OrgRecords: records}
}
