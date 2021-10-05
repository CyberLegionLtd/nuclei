package db

import (
	"embed"
	"fmt"
	"path"

	"github.com/golang-migrate/migrate/v4"
	pgx "github.com/golang-migrate/migrate/v4/database/pgx"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/pkg/errors"
)

// migrationData holds our static migration files
//go:embed migrations/*
var migrationData embed.FS

const migrationPath = "migrations"

// ApplyMigrations applies db related migrations to a postgres URL connection string
func ApplyMigrations(postgresURL string) error {
	entries, err := migrationData.ReadDir(migrationPath)
	if err != nil {
		return err
	}
	var filenames []string
	for _, entry := range entries {
		filenames = append(filenames, entry.Name())
	}

	res := bindata.Resource(filenames,
		func(name string) ([]byte, error) {
			return migrationData.ReadFile(path.Join(migrationPath, name))
		})

	postgres := &pgx.Postgres{}
	driver, err := postgres.Open(postgresURL)
	if err != nil {
		return errors.Wrap(err, "could not connect to db")
	}
	defer driver.Close()

	d, err := bindata.WithInstance(res)
	if err != nil {
		return errors.Wrap(err, "could not read migrations")
	}

	m, err := migrate.NewWithInstance("go-bindata", d, "pgx", driver)
	if err != nil {
		return fmt.Errorf("initializing db migration failed %s", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating database failed %s", err)
	}
	return nil
}
