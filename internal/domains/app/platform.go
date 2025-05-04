package app

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/hooks"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

func NewApp(info BasicInfo) Platform {
	return Platform{
		BaseHooks:     hooks.BaseHooks{},
		BaseValidator: validator.BaseValidator{},
		Auditable:     fields.NewAuditable(),
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	// ID is the unique identifier of an app Platform.
	ID int `db:"ID"`
	// ProductID is the identifier of the product that this Platform belongs.
	ProductID int `db:"ProductID"`
	// OS is the target operating system of the Platform.
	OS values.OS `db:"OS"`
}

type Platform struct {
	hooks.BaseHooks
	validator.BaseValidator
	fields.Auditable
	BasicInfo
}

// IsValid checks if the current platform details follow the pre-defined business rules
func (a *Platform) IsValid() error {
	if validator.IsAnyEmpty(string(a.OS)) {
		return ErrEmptyField
	}
	return nil
}
