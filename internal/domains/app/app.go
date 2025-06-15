package app

import (
	"database/sql"
	"errors"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

var (
	ErrEmptyField = errors.New("a required field is empty")
)

func NewApp(info BasicInfo) App {
	return App{
		BaseValidator: validator.BaseValidator{},
		Auditable:     fields.NewAuditable(),
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	// ID is the unique identifier of an App.
	ID int `xorm:"pk autoincr"`
	// Name is a human-readable unique identifier of an App.
	Name string `xorm:"varchar(100) not null unique"`
	// Slug is a URL-friendly version of the App's name.
	Slug values.Slug `xorm:"varchar(100) not null unique"`
	// Description provides details about the App.
	Description sql.NullString `xorm:"text"`
	// Private indicates whether the App is private or publicly available.
	Private bool `xorm:"not null default false"`
	// VersionFormat defines the versioning format of an App.
	VersionFormat values.VersionFormat `xorm:"varchar(20)"`
	// SetupGuideCompleted marks the starting guide as completed.
	SetupGuideCompleted bool `xorm:"not null default false"`
}

type PlatformAvailability struct {
	HasAndroid bool
	HasIOS     bool
	HasWindows bool
	HasLinux   bool
	HasMacOS   bool
}

type App struct {
	validator.BaseValidator `xorm:"-"`
	fields.Auditable        `xorm:"extends"`
	BasicInfo               `xorm:"extends"`
	PlatformAvailability    `xorm:"-"`
	// Logo is the image logo of the App.
	Logo sql.NullString
}

// IsValid checks if the current user information follows the pre-defined business rules
func (p *App) IsValid() error {
	if validator.IsAnyEmpty(p.Name) {
		return ErrEmptyField
	}
	return nil
}
