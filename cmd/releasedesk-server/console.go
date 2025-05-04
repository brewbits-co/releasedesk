package main

import (
	"github.com/brewbits-co/releasedesk/internal/controllers/auth"
	"github.com/brewbits-co/releasedesk/internal/controllers/build"
	"github.com/brewbits-co/releasedesk/internal/controllers/misc"
	"github.com/brewbits-co/releasedesk/internal/controllers/platform"
	"github.com/brewbits-co/releasedesk/internal/controllers/product"
	"github.com/brewbits-co/releasedesk/internal/controllers/release"
	authSrv "github.com/brewbits-co/releasedesk/internal/services/auth"
	"github.com/brewbits-co/releasedesk/pkg/middlewares"
	"github.com/brewbits-co/releasedesk/pkg/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"go.uber.org/dig"
	"log"
	"net/http"
)

// buildConsole sets up the Chi router with necessary middleware, resolves dependencies, configures routes, and serves static files.
func buildConsole(container *dig.Container) *http.Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/homepage", http.StatusTemporaryRedirect)
	})

	// Resolve dependencies and set up endpoints and views
	err := container.Invoke(func(
		authService authSrv.Service,
		authCtrl auth.AuthController,
		miscCtrl misc.MiscController,
		productCtrl product.ProductController,
		appCtrl platform.PlatformController,
		releaseCtrl release.ReleaseController,
		buildCtrl build.BuildController,
	) {
		// Public Routes
		router.Group(func(r chi.Router) {
			// Public Views
			r.Get("/login", authCtrl.RenderLogin)
			// Public Internal API
			r.Post("/internal/login", authCtrl.HandleLogin)
		})

		// Private Routes
		router.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(session.TokenAuth))
			r.Use(middlewares.RedirectOnUnauthorized(session.TokenAuth))

			// Private Views
			r.Get("/homepage", miscCtrl.RenderHomepage)
			r.Get("/dashboard/{slug}", productCtrl.RenderDashboard)
			r.Get("/dashboard/{slug}/releases", releaseCtrl.RenderReleaseList)
			r.Get("/dashboard/{slug}/releases/{version}", releaseCtrl.RenderReleaseSummary)
			r.Get("/dashboard/{slug}/releases/{version}/release-notes", releaseCtrl.RenderReleaseNotes)
			r.Get("/dashboard/{slug}/apps/{platform}/builds", buildCtrl.RenderBuildList)
			r.Get("/dashboard/{slug}/apps/{platform}/builds/{number}", buildCtrl.RenderBuildDetails)
			r.Get("/dashboard/{slug}/apps/{platform}/builds/{number}/metadata", buildCtrl.RenderBuildMetadata)
			// Private Internal API
			r.Post("/internal/logout", authCtrl.HandleLogout)
			r.Post("/internal/products", productCtrl.HandleCreateProduct)
			r.Post("/internal/products/{slug}/setup", productCtrl.HandleProductSetupGuide)
			r.Post("/internal/products/{slug}/apps", appCtrl.HandleAddPlatform)
			r.Post("/internal/products/{slug}/releases", releaseCtrl.HandleCreateRelease)
			r.Post("/internal/products/{slug}/apps/{platform}/builds", buildCtrl.HandleBuildUpload)
			r.Get("/internal/artifacts/{checksum}", buildCtrl.HandleArtifactDownload)
		})

		// Public API Routes
		router.Group(func(r chi.Router) {
			r.Use(middlewares.APITokenAuthorization(authService))

			r.Post("/api/products/{slug}/apps/{platform}/builds", buildCtrl.HandleBuildUpload)
		})
	})
	if err != nil {
		log.Fatalf("Failed to invoke container: %v", err)
	}

	// Handle static files from the assets folder
	router.Handle("/assets/*", http.FileServer(http.FS(assetsFS)))

	// Create HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return server
}
