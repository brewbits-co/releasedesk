package app

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/hooks"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

func NewApp(info BasicInfo) App {
	return App{
		BaseHooks:     hooks.BaseHooks{},
		BaseValidator: validator.BaseValidator{},
		Auditable:     fields.NewAuditable(),
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	// ID is the unique identifier of an App.
	ID int `db:"ID"`
	// ProductID is the identifier of the product that this App belongs.
	ProductID int `db:"ProductID"`
	// Name is a human-readable unique identifier of an App.
	Name string `db:"Name"`
	// Platform is the target OS of the App.
	Platform values.Platform `db:"Platform"`
}

type App struct {
	hooks.BaseHooks
	validator.BaseValidator
	fields.Auditable
	BasicInfo
}

// IsValid checks if the current app details follows the pre-defined business rules
func (a *App) IsValid() error {
	if validator.IsAnyEmpty(a.Name, string(a.Platform)) {
		return ErrEmptyField
	}
	return nil
}
