package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrgUserList []OrgUser

type OrgUser struct {
	Name    string         `json:"name"`
	Address sdk.AccAddress `json:"addresss"`
	Role    string         `json:"role"`
}

type Printer struct {
	Response interface{}
}

type Org struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
	Users []OrgUser      `json:"users"`
}

// implement fmt.Stringer
func (p OrgUserList) String() string {
	return fmt.Sprintf("%v", p)
}

func (p Org) String() string {
	return fmt.Sprintf("%v", p)
}

func NewOrg(name string, owner sdk.AccAddress) Org {
	return Org{
		Name:  name,
		Owner: owner,
		Users: []OrgUser{},
	}
}
