package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/dig"
	"net/http"
)

// buildPortal sets up the Chi router with necessary middleware, resolves dependencies, configures routes, and serves static files.
func buildPortal(container *dig.Container) *http.Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "Download Portal")
	})

	// Handle static files from the assets folder
	router.Handle("/assets/*", http.FileServer(http.FS(assetsFS)))

	// Create HTTP server
	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	return server
}
