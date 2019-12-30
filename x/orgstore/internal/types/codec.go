package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateOrg{}, "orgstore/CreateOrg", nil)
	cdc.RegisterConcrete(MsgDeleteOrg{}, "orgstore/DeleteOrg", nil)
	cdc.RegisterConcrete(MsgAddUser{}, "orgstore/AddUser", nil)
	cdc.RegisterConcrete(MsgDeleteUser{}, "orgstore/DeleteUser", nil)
}
