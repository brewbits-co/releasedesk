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

type CurrentProductData struct {
	ProductID   int
	ProductName string
	ProductSlug values.Slug
	Platforms   []CurrentPlatformData
}

type CurrentPlatformData struct {
	PlatformID  int
	OS          values.OS
	ProductSlug values.Slug
}

func NewCurrentPlatformData(ctx CurrentProductData, platform values.OS) CurrentPlatformData {
	for _, appCtx := range ctx.Platforms {
		if appCtx.OS == platform {
			return appCtx
		}
	}
	// Return an empty CurrentPlatformData if no match is found
	return CurrentPlatformData{}
}
