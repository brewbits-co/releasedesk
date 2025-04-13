package build

import (
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type BuildListData struct {
	session.SessionData
	session.CurrentProductData
	CurrentApp session.CurrentProductAppData
	Builds     []build.BasicInfo
}

type BuildDetailsData struct {
	session.SessionData
	session.CurrentProductData
	CurrentApp session.CurrentProductAppData
	Build      build.Build
}
