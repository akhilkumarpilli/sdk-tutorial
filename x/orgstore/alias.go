package orgstore

import (
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/keeper"
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgCreateOrg  = types.NewMsgCreateOrg
	NewMsgDeleteOrg  = types.NewMsgDeleteOrg
	NewMsgAddUser    = types.NewMsgAddUser
	NewMsgDeleteUser = types.NewMsgDeleteUser
	NewOrg           = types.NewOrg
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper        = keeper.Keeper
	MsgCreateOrg  = types.MsgCreateOrg
	MsgDeleteOrg  = types.MsgDeleteOrg
	MsgAddUser    = types.MsgAddUser
	MsgDeleteUser = types.MsgDeleteUser
	QueryResOrgs  = types.QueryResOrgs
	Org           = types.Org
)
