package main

import (
	"database/sql"
	"github.com/brewbits-co/releasedesk/internal/domains/app"
	"github.com/brewbits-co/releasedesk/internal/domains/build"
	"github.com/brewbits-co/releasedesk/internal/domains/platform"
	"github.com/brewbits-co/releasedesk/internal/domains/release"
	"github.com/brewbits-co/releasedesk/internal/domains/user"
	"github.com/brewbits-co/releasedesk/internal/values"
	"github.com/brewbits-co/releasedesk/pkg/fields"
	"github.com/brewbits-co/releasedesk/pkg/validator"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"
)

func applyMigrations(engine *xorm.Engine) error {
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
		{
			ID: "202506142204",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&platform.Platform{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&platform.Platform{})
			},
		},
		{
			ID: "202506142205",
			Migrate: func(tx *xorm.Engine) error {
				err := tx.Sync2(&user.User{})
				if err != nil {
					return err
				}
				_, err = tx.Insert(user.User{
					BaseValidator: validator.BaseValidator{},
					Auditable:     fields.Auditable{},
					ID:            0,
					Username:      "admin",
					Email:         sql.NullString{},
					Password:      "$2a$10$Z13RQlu6HdKSW41rJsz7Ju5NZ0VMyUdm6YZMr0wjJqW955qd2pzx2", // admin
					FirstName:     sql.NullString{},
					LastName:      sql.NullString{},
					Role:          values.Admin,
				})
				if err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&user.User{})
			},
		},
		{
			ID: "202506142206",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&release.Release{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&release.Release{})
			},
		},
		{
			ID: "202506142207",
			Migrate: func(tx *xorm.Engine) error {
				type Release struct {
					ReleaseNotes string `xorm:"longtext"`
				}

				return tx.Sync2(&release.Changelog{}, &Release{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&release.Changelog{})
			},
		},
		{
			ID: "202506162330",
			Migrate: func(tx *xorm.Engine) error {
				type Release struct {
					BuildSelection string `xorm:"varchar(10) not null default 'Last'"`
				}

				// Create LinkedBuilds table
				err := tx.Sync2(&release.LinkedBuilds{}, &Release{})
				if err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *xorm.Engine) error {
				// Drop LinkedBuilds table
				err := tx.DropTables(&release.LinkedBuilds{})
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			ID: "202506191930",
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync2(&release.ChecklistItem{})
			},
			Rollback: func(tx *xorm.Engine) error {
				return tx.DropTables(&release.ChecklistItem{})
			},
		},
	}

	m := migrate.New(engine, &migrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}, migrations)

	return m.Migrate()
}
