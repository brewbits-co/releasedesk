package release

import (
	"github.com/brewbits-co/releasedesk/pkg/hooks"
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
	ID int `db:"ID"`
	// ReleaseID is the identifier of the release that this changelog belongs to
	ReleaseID int `db:"ReleaseID"`
	// Text is the description of the change
	Text string `db:"Text"`
	// ChangeType indicates the type of change
	ChangeType ChangeType `db:"ChangeType"`
}

// NewReleaseNotes creates a new ReleaseNotes entity
func NewReleaseNotes(releaseID int, text string) ReleaseNotes {
	return ReleaseNotes{
		BaseHooks:     hooks.BaseHooks{},
		BaseValidator: validator.BaseValidator{},
		ReleaseID:     releaseID,
		Text:          text,
		Changelogs:    []Changelog{},
	}
}

// ReleaseNotes represents the release notes for a specific release
type ReleaseNotes struct {
	hooks.BaseHooks
	validator.BaseValidator
	// ID is the unique identifier of the ReleaseNotes
	ID int `db:"ID"`
	// ReleaseID is the identifier of the release that these notes belong to
	ReleaseID int `db:"ReleaseID"`
	// Text is the general description or overview of the release
	Text string `db:"Text"`
	// Changelogs is a list of changelog entries for this release
	Changelogs []Changelog `db:"-"`
}
