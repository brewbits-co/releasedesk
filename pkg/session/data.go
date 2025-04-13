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
	Apps        []CurrentProductAppData
}

type CurrentProductAppData struct {
	AppID       int
	AppName     string
	AppPlatform values.Platform
	ProductSlug values.Slug
}

func NewCurrentAppData(ctx CurrentProductData, platform values.Platform) CurrentProductAppData {
	for _, appCtx := range ctx.Apps {
		if appCtx.AppPlatform == platform {
			return appCtx
		}
	}
	// Return an empty CurrentProductAppData if no match is found
	return CurrentProductAppData{}
}
