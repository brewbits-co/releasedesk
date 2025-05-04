package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type ReleaseListData struct {
	session.SessionData
	session.CurrentProductData
	Releases         []release.BasicInfo
	Channels         []release.Channel
	CurrentChannelID int
}

type ReleaseSummaryData struct {
	session.SessionData
	session.CurrentProductData
	release.Release
	Channels []release.Channel
}

type ReleaseNotesData struct {
	session.SessionData
	session.CurrentProductData
}
