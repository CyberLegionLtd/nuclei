package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDBMigrations(t *testing.T) {
	err := ApplyMigrations(getPostgresConnString())
	require.Nil(t, err, "could not run and apply migrations to db")
}
