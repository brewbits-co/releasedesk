package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

type ReleaseStatus string

const (
	Draft       ReleaseStatus = "Draft"
	Published   ReleaseStatus = "Published"
	Deprecated  ReleaseStatus = "Deprecated"
	Unpublished ReleaseStatus = "Unpublished"
	Scheduled   ReleaseStatus = "Scheduled"
)

func NewRelease(info BasicInfo) Release {
	info.Auditable = fields.NewAuditable()
	return Release{
		BaseValidator: validator.BaseValidator{},
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	fields.Auditable `xorm:"extends"`
	// ID is the unique identifier of a Release.
	ID int `xorm:"pk autoincr"`
	// AppID is the identifier of the app that this Release belongs.
	AppID int `xorm:"not null unique(release_uidx)"`
	// Version specifies the version of the Release.
	Version string `xorm:"not null unique(release_uidx)"`
	// TargetChannel
	TargetChannel int `xorm:"not null"`
	// Status
	Status ReleaseStatus `xorm:"not null"`
}

type Release struct {
	validator.BaseValidator `xorm:"-"`
	BasicInfo               `xorm:"extends"`
	Builds                  []build.BasicInfo `xorm:"-"`
}
