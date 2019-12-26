package keeper

import (
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey,
		cdc,
	}
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Gets the entire org metadata struct for a name
func (k Keeper) GetOrgData(ctx sdk.Context, name string) types.Org {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name) {
		return types.Org{}
	}
	bz := store.Get([]byte(name))
	var org types.Org
	k.cdc.MustUnmarshalBinaryBare(bz, &org)
	return org
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetOrgData(ctx, name).Owner
}

func (k Keeper) SetOrg(ctx sdk.Context, name string, org types.Org) {
	if org.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(org))
}

// Deletes the entire org metadata struct for a name
func (k Keeper) DeleteOrg(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// Add a user to org metadata
func (k Keeper) AddUser(ctx sdk.Context, name string, msg types.MsgAddUser) {
	org := k.GetOrgData(ctx, name)
	user := types.OrgUser{
		Name:    msg.Username,
		Address: msg.UserAddr,
		Role:    msg.UserRole,
	}
	org.Users = append(org.Users, user)
	k.SetOrg(ctx, name, org)
}

// Add a user to org metadata
func (k Keeper) DeleteUser(ctx sdk.Context, name string, msg types.MsgDeleteUser) {
	org := k.GetOrgData(ctx, name)
	var orgUsers []types.OrgUser
	for _, u := range org.Users {
		if !u.Address.Equals(msg.UserAddr) {
			orgUsers = append(orgUsers, u)
		}
	}
	org.Users = orgUsers
	k.SetOrg(ctx, name, org)
}

// Get an iterator over all names in which the keys are the orgs and the values are the org data
func (k Keeper) GetOrgsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
