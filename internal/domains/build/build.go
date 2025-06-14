package build

import (
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

func NewBuild(info BasicInfo) Build {
	info.Auditable = fields.NewAuditable()
	return Build{
		BaseValidator: validator.BaseValidator{},
		BasicInfo:     info,
		Artifacts:     make([]Artifact, 0),
	}
}

type BasicInfo struct {
	fields.Auditable `xorm:"extends"`
	// ID is the unique identifier of a Build.
	ID int `xorm:"pk autoincr"`
	// PlatformID is the identifier of the platform that this Build belongs.
	PlatformID int `xorm:"not null unique(build_uidx)"`
	// Version specifies the version of the Build.
	Version string `xorm:"not null"`
	// Number represents a unique code or sequence number associated with the Build.
	Number string `xorm:"not null unique(build_uidx)"`
}

type Build struct {
	validator.BaseValidator `xorm:"-"`
	fields.Extendable       `xorm:"-"`
	BasicInfo               `xorm:"extends"`
	Artifacts               []Artifact `xorm:"-"`
}

// BuildMetadata represents metadata information associated with a build.
// Each entry is primarily used to store and manage the key-value pairs found in fields.Extendable field of a Build.
type BuildMetadata struct {
	BuildID int    `xorm:"pk"`
	Key     string `xorm:"pk"`
	Value   string
}
