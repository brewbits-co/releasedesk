package release

import (
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

// ChangeType represents the type of change in a changelog entry
type ChangeType string

const (
	// Change types
	Changed          ChangeType = "Changed"
	Added            ChangeType = "Added"
	Fixed            ChangeType = "Fixed"
	Removed          ChangeType = "Removed"
	Security         ChangeType = "Security"
	DeprecatedChange ChangeType = "Deprecated"
)

// Changelog represents a single changelog entry in release notes
type Changelog struct {
	// ID is the unique identifier of a Changelog entry
	ID int `xorm:"pk autoincr"`
	// ReleaseID is the identifier of the release that this changelog belongs to
	ReleaseID int `xorm:"not null"`
	// Text is the description of the change
	Text string `xorm:"not null"`
	// ChangeType indicates the type of change
	ChangeType ChangeType `xorm:"not null"`
}

// NewReleaseNotes creates a new ReleaseNotes entity
func NewReleaseNotes(releaseID int, text string) ReleaseNotes {
	return ReleaseNotes{
		BaseValidator: validator.BaseValidator{},
		ReleaseID:     releaseID,
		Text:          text,
		Changelogs:    []Changelog{},
	}
}

// ReleaseNotes represents the release notes for a specific release
type ReleaseNotes struct {
	validator.BaseValidator `xorm:"-"`
	// ReleaseID is the identifier of the release that these notes belong to
	ReleaseID int
	// Text is the general description or overview of the release
	Text string
	// Changelogs is a list of changelog entries for this release
	Changelogs []Changelog `xorm:"-"`
}
