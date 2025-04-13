package product

import (
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type DashboardData struct {
	session.SessionData
	session.CurrentProductData
	SetupGuideCompleted bool
}
