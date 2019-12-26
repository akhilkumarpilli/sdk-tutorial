package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgCreateOrg struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgCreateOrg(name string, owner sdk.AccAddress) MsgCreateOrg {
	return MsgCreateOrg{
		Name:  name,
		Owner: owner,
	}
}

func (msg MsgCreateOrg) Type() string {
	return "create_org"
}

func (msg MsgCreateOrg) Route() string {
	return RouterKey
}

func (msg MsgCreateOrg) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name can't be empty")
	}

	return nil
}

func (msg MsgCreateOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCreateOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgDeleteOrg struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgDeleteOrg(name string, owner sdk.AccAddress) MsgDeleteOrg {
	return MsgDeleteOrg{
		Name:  name,
		Owner: owner,
	}
}

func (msg MsgDeleteOrg) Type() string {
	return "delete_org"
}

func (msg MsgDeleteOrg) Route() string {
	return RouterKey
}

func (msg MsgDeleteOrg) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name can't be empty")
	}

	return nil
}

func (msg MsgDeleteOrg) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteOrg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgAddUser struct {
	OrgName  string         `json:"name"`
	OrgOwner sdk.AccAddress `json:"owner"`
	Username string         `json:"name"`
	UserAddr sdk.AccAddress `json:"address"`
	UserRole string         `json:"role"`
}

func NewMsgAddUser(orgName string, orgOwner sdk.AccAddress, username string, userAddr sdk.AccAddress, userRole string) MsgAddUser {
	return MsgAddUser{
		OrgName:  orgName,
		OrgOwner: orgOwner,
		UserAddr: userAddr,
		Username: username,
		UserRole: userRole,
	}
}

func (msg MsgAddUser) Type() string {
	return "add_user"
}

func (msg MsgAddUser) Route() string {
	return RouterKey
}

func (msg MsgAddUser) ValidateBasic() sdk.Error {
	if msg.OrgOwner.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf(`Invalid org address: %s`, msg.OrgOwner.String()))
	}

	if msg.UserAddr.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf(`Invalid user address: %s`, msg.UserAddr.String()))
	}

	if len(msg.OrgName) == 0 || len(msg.Username) == 0 {
		return sdk.ErrUnknownRequest("Org name or User name can't be empty")
	}

	return nil
}

func (msg MsgAddUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAddUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.OrgOwner}
}

type MsgDeleteUser struct {
	OrgName  string         `json:"name"`
	OrgOwner sdk.AccAddress `json:"owner"`
	UserAddr sdk.AccAddress `json:"address"`
}

func NewMsgDeleteUser(orgName string, orgOwner sdk.AccAddress, userAddr sdk.AccAddress) MsgDeleteUser {
	return MsgDeleteUser{
		OrgName:  orgName,
		OrgOwner: orgOwner,
		UserAddr: userAddr,
	}
}

func (msg MsgDeleteUser) Type() string {
	return "del_user"
}

func (msg MsgDeleteUser) Route() string {
	return RouterKey
}

func (msg MsgDeleteUser) ValidateBasic() sdk.Error {
	if msg.OrgOwner.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf(`Invalid org address: %s`, msg.OrgOwner.String()))
	}

	if msg.UserAddr.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf(`Invalid user address: %s`, msg.UserAddr.String()))
	}

	if len(msg.OrgName) == 0 {
		return sdk.ErrUnknownRequest("Org name or User name can't be empty")
	}

	return nil
}

func (msg MsgDeleteUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgDeleteUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.OrgOwner}
}
