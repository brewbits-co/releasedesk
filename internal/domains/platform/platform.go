package platform

import (
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

func NewPlatform(info BasicInfo) Platform {
	return Platform{
		BaseValidator: validator.BaseValidator{},
		Auditable:     fields.NewAuditable(),
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	// ID is the unique identifier of a Platform.
	ID int `xorm:"pk autoincr"`
	// AppID is the identifier of the application that this Platform belongs to.
	AppID int `xorm:"not null unique(platform_uidx)"`
	// OS is the target operating system of the Platform.
	OS values.OS `xorm:"not null unique(platform_uidx)"`
}

type Platform struct {
	validator.BaseValidator `xorm:"-"`
	fields.Auditable        `xorm:"extends"`
	BasicInfo               `xorm:"extends"`
}

// IsValid checks if the current platform details follow the pre-defined business rules
func (a *Platform) IsValid() error {
	if validator.IsAnyEmpty(string(a.OS)) {
		return ErrEmptyField
	}
	return nil
}
