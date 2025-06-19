package release

import (
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
)

// NewChecklistItem creates a new ChecklistItem entity
func NewChecklistItem(releaseID int, text string) ChecklistItem {
	return ChecklistItem{
		BaseValidator: validator.BaseValidator{},
		ReleaseID:     releaseID,
		Text:          text,
		Checked:       false,
		Orderable:     fields.Orderable{},
	}
}

// ChecklistItem represents an item in a release checklist
type ChecklistItem struct {
	validator.BaseValidator `xorm:"-"`
	// ID is the unique identifier of a ChecklistItem
	ID int `xorm:"pk autoincr"`
	// ReleaseID is the identifier of the release that this checklist item belongs to
	ReleaseID int `xorm:"not null index"`
	// Text is the description of the checklist item
	Text string `xorm:"not null"`
	// Checked indicates whether the checklist item has been completed
	Checked bool `xorm:"not null default false"`
	// Orderable provides ordering functionality
	fields.Orderable `xorm:"extends"`
}
