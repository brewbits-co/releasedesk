package main

import (
	"embed"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

// Include the sql files in the binary
//
//go:embed migrations/*.sql
var migrationFiles embed.FS

// Initialize the SQLite connection and return *sql.DB
func newSQLiteDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./_data/releasedesk.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create a migration source using the embedded files
	d, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		log.Fatalf("Could not create iofs driver: %v", err)
	}

	// Configure the SQLite migration driver
	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Could not create SQLite driver instance: %v", err)
	}

	// Initialize migration with an embedded source and SQLite driver
	m, err := migrate.NewWithInstance(
		"iofs", d, "sqlite3", driver,
	)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	// Enable WAL mode for better concurrency
	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	log.Println("Migrations applied successfully!")

	return db, nil
}
