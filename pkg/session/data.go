package session

import (
	"context"
	"github.com/brewbits-co/releasedesk/internal/values"
)

type SessionData struct {
	Username string
}

func NewSessionData(ctx context.Context) SessionData {
	return SessionData{
		Username: ctx.Value("username").(string),
	}
}

type CurrentAppData struct {
	AppID     int
	AppName   string
	AppSlug   values.Slug
	Platforms []CurrentPlatformData
}

type CurrentPlatformData struct {
	PlatformID int
	OS         values.OS
	AppSlug    values.Slug
}

func NewCurrentPlatformData(ctx CurrentAppData, platform values.OS) CurrentPlatformData {
	for _, appCtx := range ctx.Platforms {
		if appCtx.OS == platform {
			return appCtx
		}
	}
	// Return an empty CurrentPlatformData if no match is found
	return CurrentPlatformData{}
}
