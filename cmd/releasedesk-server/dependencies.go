package main

import (
	"embed"
	appCtrl "github.com/brewbits-co/releasedesk/internal/controllers/app"
	authCtrl "github.com/brewbits-co/releasedesk/internal/controllers/auth"
	buildCtrl "github.com/brewbits-co/releasedesk/internal/controllers/build"
	"github.com/brewbits-co/releasedesk/internal/controllers/misc"
	platformCtrl "github.com/brewbits-co/releasedesk/internal/controllers/platform"
	releaseCtrl "github.com/brewbits-co/releasedesk/internal/controllers/release"
	"github.com/brewbits-co/releasedesk/internal/persistence/sql"
	"github.com/brewbits-co/releasedesk/internal/services/app"
	"github.com/brewbits-co/releasedesk/internal/services/auth"
	"github.com/brewbits-co/releasedesk/internal/services/build"
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"github.com/brewbits-co/releasedesk/internal/services/release"
	"go.uber.org/dig"
)

// Include the assets in the binary
//
//go:embed assets/*
var assetsFS embed.FS

// Register repository, service, and controller in Uber Dig container
func buildContainer() *dig.Container {
	container := dig.New()

	// Register the constructors using factory methods
	container.Provide(newDBEngine)
	// Repositories
	container.Provide(sql.NewUserRepository)
	container.Provide(sql.NewApplicationRepository)
	container.Provide(sql.NewPlatformRepository)
	container.Provide(sql.NewBuildRepository)
	container.Provide(sql.NewReleaseRepository)
	container.Provide(sql.NewReleaseNotesRepository)
	container.Provide(sql.NewChecklistRepository)
	// Services
	container.Provide(auth.NewAuthService)
	container.Provide(app.NewAppService)
	container.Provide(platform.NewPlatformService)
	container.Provide(build.NewBuildService)
	container.Provide(release.NewReleaseService)
	// Controllers
	container.Provide(authCtrl.NewAuthController)
	container.Provide(misc.NewMiscController)
	container.Provide(appCtrl.NewAppController)
	container.Provide(releaseCtrl.NewReleaseController)
	container.Provide(platformCtrl.NewPlatformController)
	container.Provide(buildCtrl.NewBuildController)

	return container
}
