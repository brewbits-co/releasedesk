package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/values"
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

type BuildSelection string

const (
	Last   BuildSelection = "Last"
	Manual BuildSelection = "Manual"
)

func NewRelease(info BasicInfo) Release {
	info.Auditable = fields.NewAuditable()
	return Release{
		BaseValidator: validator.BaseValidator{},
		BasicInfo:     info,
		Builds:        make(map[values.OS]build.BasicInfo),
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
	// BuildSelection specifies how builds are selected for this release
	BuildSelection BuildSelection `xorm:"not null default 'Last'"`
}

// LinkedBuilds represents the relationship between a Release and a Build
type LinkedBuilds struct {
	// ReleaseID is the identifier of the release
	ReleaseID int `xorm:"pk not null"`
	// BuildID is the identifier of the build
	BuildID int `xorm:"pk not null"`
	// OS is the operating system that this build is for
	OS values.OS `xorm:"pk not null"`
}

type Release struct {
	validator.BaseValidator `xorm:"-"`
	BasicInfo               `xorm:"extends"`
	Builds                  map[values.OS]build.BasicInfo `xorm:"-"`
}
