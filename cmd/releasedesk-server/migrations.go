package main

import (
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"
)

func applyMigrations(engine *xorm.Engine, err error) {
	var migrations = []*migrate.Migration{
		{
			ID: "202506142200",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&app.App{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&app.App{})
			},
		},
		{
			ID: "202506142201",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&release.Channel{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&release.Channel{})
			},
		},
		{
			ID: "202506142202",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&build.Build{}, &build.BuildMetadata{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&build.Build{}, &build.BuildMetadata{})
			},
		},
		{
			ID: "202506142203",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&build.Artifact{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&build.Build{})
			},
		},
	}

	m := migrate.New(engine, &migrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}, migrations)

	err = m.Migrate()
}
