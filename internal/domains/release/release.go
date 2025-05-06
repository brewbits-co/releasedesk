package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/hooks"
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
		BaseHooks:     hooks.BaseHooks{},
		BaseValidator: validator.BaseValidator{},
		BasicInfo:     info,
	}
}

type BasicInfo struct {
	fields.Auditable
	// ID is the unique identifier of a Release.
	ID int `db:"ID"`
	// AppID is the identifier of the app that this Release belongs.
	AppID int `db:"AppID"`
	// Version specifies the version of the Release.
	Version string `db:"Version"`
	// TargetChannel
	TargetChannel int `db:"TargetChannel"`
	// Status
	Status ReleaseStatus `db:"Status"`
}

type Release struct {
	hooks.BaseHooks
	validator.BaseValidator
	BasicInfo
	Builds []build.BasicInfo `db:"-"`
}
