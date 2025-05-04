package main

import (
	"embed"
	authCtrl "github.com/brewbits-co/releasedesk/internal/controllers/auth"
	buildCtrl "github.com/brewbits-co/releasedesk/internal/controllers/build"
	"github.com/brewbits-co/releasedesk/internal/controllers/misc"
	appCtrl "github.com/brewbits-co/releasedesk/internal/controllers/platform"
	productCtrl "github.com/brewbits-co/releasedesk/internal/controllers/product"
	releaseCtrl "github.com/brewbits-co/releasedesk/internal/controllers/release"
	"github.com/brewbits-co/releasedesk/internal/persistence/sql"
	"github.com/brewbits-co/releasedesk/internal/services/auth"
	"github.com/brewbits-co/releasedesk/internal/services/build"
	"github.com/brewbits-co/releasedesk/internal/services/platform"
	"github.com/brewbits-co/releasedesk/internal/services/product"
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
	container.Provide(newSQLiteDB)
	// Repositories
	container.Provide(sql.NewUserRepository)
	container.Provide(sql.NewProductRepository)
	container.Provide(sql.NewPlatformRepository)
	container.Provide(sql.NewBuildRepository)
	container.Provide(sql.NewReleaseRepository)
	// Services
	container.Provide(auth.NewAuthService)
	container.Provide(product.NewProductService)
	container.Provide(platform.NewPlatformService)
	container.Provide(build.NewBuildService)
	container.Provide(release.NewReleaseService)
	// Controllers
	container.Provide(authCtrl.NewAuthController)
	container.Provide(misc.NewMiscController)
	container.Provide(productCtrl.NewProductController)
	container.Provide(releaseCtrl.NewReleaseController)
	container.Provide(appCtrl.NewPlatformController)
	container.Provide(buildCtrl.NewBuildController)

	return container
}
