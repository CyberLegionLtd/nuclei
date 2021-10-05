package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
	"github.com/stretchr/testify/require"
)

func TestTemplatesService(t *testing.T) {
	conn, err := New(getPostgresConnString())
	require.Nil(t, err, "could not get db connection")

	templates := newTemplatesService(conn.DB())

	addition := &model.Template{
		ID:         uuid.New().String(),
		Path:       "test/template.yaml",
		Name:       "Test Template",
		IsWorkflow: false,
		Contents:   "test contents",
	}
	t.Run("add", func(t *testing.T) {
		err := templates.Add(addition)
		require.Nil(t, err, "could not add template to db")
	})

	var got *model.Template
	t.Run("list", func(t *testing.T) {
		err := templates.List(func(template *model.Template) {
			got = template
		})
		require.Nil(t, err, "could not list templates")
		require.NotNil(t, got, "could not list templates")
		require.Equal(t, addition.Path, got.Path, "could not list templates correctly")
	})

	t.Run("get", func(t *testing.T) {
		contents, err := templates.Get(got.ID)
		require.Nil(t, err, "could not get template contents")
		require.Equal(t, addition.Contents, contents, "could not get template contents")
	})

	t.Run("update", func(t *testing.T) {
		addition.Contents = "updated contents"
		err = templates.Add(addition)
		require.Nil(t, err, "could not update template contents")

		contents, err := templates.Get(got.ID)
		require.Nil(t, err, "could not get updated template contents")
		require.Equal(t, addition.Contents, contents, "could not get updated template contents")
	})

	t.Run("delete", func(t *testing.T) {
		err = templates.Delete(got.ID)
		require.Nil(t, err, "could not delete template id")

		_, err = templates.Get(got.ID)
		require.NotNil(t, err, "could get deleted template contents")
	})
}
