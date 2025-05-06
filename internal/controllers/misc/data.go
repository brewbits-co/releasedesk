package misc

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type HomepageData struct {
	session.SessionData
	Apps []app.App
}
