package build

import (
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/hooks"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

func NewBuild(info BasicInfo) Build {
	info.Auditable = fields.NewAuditable()
	return Build{
		BaseHooks:     hooks.BaseHooks{},
		BaseValidator: validator.BaseValidator{},
		BasicInfo:     info,
		Artifacts:     make([]Artifact, 0),
	}
}

type BasicInfo struct {
	fields.Auditable
	// ID is the unique identifier of a Build.
	ID int `db:"ID"`
	// PlatformID is the identifier of the platform that this Build belongs.
	PlatformID int `db:"PlatformID"`
	// Version specifies the version of the Build.
	Version string `db:"Version"`
	// Number represents a unique code or sequence number associated with the Build.
	Number string `db:"Number"`
}

type Build struct {
	hooks.BaseHooks
	validator.BaseValidator
	fields.Extendable
	BasicInfo
	Artifacts []Artifact `db:"-"`
}
