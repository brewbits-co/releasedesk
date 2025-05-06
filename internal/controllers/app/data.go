package app

import (
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type DashboardData struct {
	session.SessionData
	session.CurrentAppData
	SetupGuideCompleted bool
}
