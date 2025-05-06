package misc

import (
	"github.com/brewbits-co/releasedesk/internal/domains/product"
	"github.com/brewbits-co/releasedesk/pkg/session"
)

type HomepageData struct {
	session.SessionData
	Apps []product.App
}
