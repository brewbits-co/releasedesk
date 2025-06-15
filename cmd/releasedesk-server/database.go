package main

import (
	"log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the SQLite connection and return *xorm.Engine
func newDBEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "./_data/releasedesk.db")
	if err != nil {
		log.Fatal(err)
	}

	engine.SetMapper(names.GonicMapper{})
	engine.ShowSQL(true)

	applyMigrations(engine, err)

	return engine, nil
}
