package orgstore

import (
	"fmt"
	"github.com/akhilkumarpilli/sdk-tutorial/x/orgstore/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateOrg:
			return handleMsgCreateOrg(ctx, keeper, msg)
		case MsgDeleteOrg:
			return handleMsgDeleteOrg(ctx, keeper, msg)
		case MsgAddUser:
			return handleMsgAddUser(ctx, keeper, msg)
		case MsgDeleteUser:
			return handleMsgDelUser(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized orgstore Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to create org
func handleMsgCreateOrg(ctx sdk.Context, keeper Keeper, msg MsgCreateOrg) sdk.Result {
	if keeper.IsNamePresent(ctx, msg.Name) { // Checks if the the org name already exists or not
		return types.ErrNameAlreadyExist(types.DefaultCodeSpace).Result()
	}
	org := types.Org{
		Name:  msg.Name,
		Owner: msg.Owner,
		Users: []types.OrgUser{},
	}
	keeper.SetOrg(ctx, msg.Name, org) // If so, set the name to the value specified in the msg.
	return sdk.Result{}
}

// Handle a message to delete org
func handleMsgDeleteOrg(ctx sdk.Context, keeper Keeper, msg MsgDeleteOrg) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.Name) { // Checks if the the org name already exists or not
		return types.ErrNameDoesNotExist(types.DefaultCodeSpace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteOrg(ctx, msg.Name)
	return sdk.Result{}
}

// Handle a message to add user
func handleMsgAddUser(ctx sdk.Context, keeper Keeper, msg MsgAddUser) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.OrgName) { // Checks if the the org name already exists or not
		return types.ErrNameDoesNotExist(types.DefaultCodeSpace).Result()
	}
	if !msg.OrgOwner.Equals(keeper.GetOwner(ctx, msg.OrgName)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.AddUser(ctx, msg.OrgName, msg)
	return sdk.Result{}
}

// Handle a message to delete user
func handleMsgDelUser(ctx sdk.Context, keeper Keeper, msg MsgDeleteUser) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.OrgName) { // Checks if the the org name already exists or not
		return types.ErrNameDoesNotExist(types.DefaultCodeSpace).Result()
	}
	if !msg.OrgOwner.Equals(keeper.GetOwner(ctx, msg.OrgName)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteUser(ctx, msg.OrgName, msg)
	return sdk.Result{}
}
