package fields

import (
	"github.com/brewbits-co/releasedesk/pkg/utils"
	"time"
)

func NewAuditable() Auditable {
	now := time.Now()
	nowFormatted := utils.FormatTime(now)
	return Auditable{
		CreatedAt:          now,
		CreatedAtFormatted: nowFormatted,
		UpdatedAt:          now,
		UpdatedAtFormatted: nowFormatted,
	}
}

// Auditable struct holds common timestamp and formatted fields for creation and updates
type Auditable struct {
	// CreatedAt is the timestamp when the entity was created.
	CreatedAt time.Time `db:"CreatedAt" xorm:"created"`
	// CreatedAtFormatted is a human-readable version of CreatedAt.
	CreatedAtFormatted string `xorm:"-"`
	// UpdatedAt is the timestamp when the entity was last updated.
	UpdatedAt time.Time `db:"UpdatedAt" xorm:"updated"`
	// UpdatedAtFormatted is a human-readable version of UpdatedAt.
	UpdatedAtFormatted string `xorm:"-"`
}

func (a *Auditable) FormatAuditable() {
	a.CreatedAtFormatted = utils.FormatTime(a.CreatedAt)
	a.UpdatedAtFormatted = utils.FormatTime(a.UpdatedAt)
}
