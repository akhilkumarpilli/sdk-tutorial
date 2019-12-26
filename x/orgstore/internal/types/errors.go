package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodeSpace is the Module Name
const (
	DefaultCodeSpace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist sdk.CodeType = 101
	CodeNameAlreadyExist sdk.CodeType = 102
)

// ErrNameDoesNotExist is the error for name not existing
func ErrNameDoesNotExist(codeSpace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codeSpace, CodeNameDoesNotExist, "Name does not exist")
}

// ErrNameAlreadyExist is the error for name already existing
func ErrNameAlreadyExist(codeSpace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codeSpace, CodeNameAlreadyExist, "Name already exist")
}
