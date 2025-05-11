package release

import (
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type ReleaseListData struct {
	CurrentPage string
	session.SessionData
	session.CurrentAppData
	Releases         []release.BasicInfo
	Channels         []release.Channel
	CurrentChannelID int
}

type ReleaseSummaryData struct {
	CurrentPage string
	session.SessionData
	session.CurrentAppData
	release.Release
	Channels []release.Channel
}

type ReleaseNotesData struct {
	CurrentPage string
	session.SessionData
	session.CurrentAppData
	Release      release.Release
	ReleaseNotes release.ReleaseNotes
}
