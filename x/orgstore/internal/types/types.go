package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
func (p Printer) String() string {
	return fmt.Sprintf("%s", p)
}

func NewOrg(name string, owner sdk.AccAddress) Org {
	return Org{
		Name:  name,
		Owner: owner,
		Users: []OrgUser{},
	}
}
