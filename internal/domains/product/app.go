package product

import (
	"database/sql"
	"errors"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/hooks"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

var (
	ErrEmptyField = errors.New("a required field is empty")
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
	// Name is a human-readable unique identifier of an App.
	Name string `db:"Name"`
	// Slug is a URL-friendly version of the App's name.
	Slug values.Slug `db:"Slug"`
	// Description provides details about the App.
	Description sql.NullString `db:"Description"`
	// Private indicates whether the App is private or publicly available.
	Private bool `db:"Private"`
	// VersionFormat defines the versioning format of an App.
	VersionFormat values.VersionFormat `db:"VersionFormat"`
	// SetupGuideCompleted marks the starting guide as completed.
	SetupGuideCompleted bool `db:"SetupGuideCompleted"`
}

type PlatformAvailability struct {
	HasAndroid bool `db:"HasAndroid"`
	HasIOS     bool `db:"HasIOS"`
	HasWindows bool `db:"HasWindows"`
	HasLinux   bool `db:"HasLinux"`
	HasMacOS   bool `db:"HasMacOS"`
}

type App struct {
	hooks.BaseHooks
	validator.BaseValidator
	fields.Auditable
	BasicInfo
	PlatformAvailability
	// Logo is the image logo of the App.
	Logo sql.NullString `db:"Logo"`
}

// IsValid checks if the current user information follows the pre-defined business rules
func (p *App) IsValid() error {
	if validator.IsAnyEmpty(p.Name) {
		return ErrEmptyField
	}
	return nil
}
