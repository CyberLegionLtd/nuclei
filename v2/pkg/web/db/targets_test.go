package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
	"github.com/stretchr/testify/require"
)

func TestTargetsService(t *testing.T) {
	conn, err := New(getPostgresConnString())
	require.Nil(t, err, "could not get db connection")

	targets := newTargetsService(conn.DB())

	addition := &model.Target{
		ID:         uuid.New().String(),
		Name:       "Test Target",
		RawPath:    "test",
		TotalHosts: 100,
	}
	t.Run("add", func(t *testing.T) {
		err := targets.Add(addition)
		require.Nil(t, err, "could not add target to db")
	})

	var got *model.Target
	t.Run("list", func(t *testing.T) {
		err := targets.List(func(target *model.Target) {
			got = target
		})
		require.Nil(t, err, "could not list targets")
		require.NotNil(t, got, "could not list targets")
		require.Equal(t, addition.RawPath, got.RawPath, "could not list targets correctly")
	})

	t.Run("delete", func(t *testing.T) {
		err = targets.Delete(got.ID)
		require.Nil(t, err, "could not delete targets id")
	})
}
