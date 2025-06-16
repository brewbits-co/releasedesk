package main

import (
	"log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the SQLite connection and return *xorm.Engine
func newDBEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "./_data/releasedesk.db?_pragma=journal_mode(WAL)")
	if err != nil {
		log.Fatalf("Failed to start database engine: %v", err)
	}

	// Ping the database to verify the connection is alive
	err = engine.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	engine.SetMapper(names.GonicMapper{})
	engine.ShowSQL(true)

	err = applyMigrations(engine)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Enable WAL mode for better concurrency
	_, err = engine.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	return engine, nil
}
