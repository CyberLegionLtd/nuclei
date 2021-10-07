package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
)

type templatesService struct {
	db *pgxpool.Pool
}

func newTemplatesService(db *pgxpool.Pool) *templatesService {
	return &templatesService{db: db}
}

// List returns all templates in db without their content
func (s *templatesService) List(callback func(template *model.Template)) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, `SELECT
id,
"isWorkflow",
name,
path,
"createdAt",
"updatedAt" FROM templates;`)
	if err != nil {
		return errors.Wrap(err, "could not list templates")
	}
	defer rows.Close()

	template := &model.Template{}
	for rows.Next() {
		if scanErr := rows.Scan(&template.ID, &template.IsWorkflow, &template.Name, &template.Path, &template.CreatedAt, &template.UpdatedAt); scanErr != nil {
			err = scanErr
		}
		callback(template)
	}
	return err
}

// Get returns the contents for a template ID
func (s *templatesService) Get(ID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var contents string
	err := s.db.QueryRow(ctx, `SELECT
contents
FROM templates WHERE ID=$1;`, ID).Scan(&contents)
	if err != nil {
		return "", errors.Wrap(err, "could not get template content")
	}
	return contents, nil
}

// Add adds a template to the db.
// If the template already exists for a path, it's contents are updated.
func (s *templatesService) Add(template *model.Template) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.db.Exec(ctx, `INSERT INTO templates (
id,
"isWorkflow",
name,
path,
contents,
"createdAt",
"updatedAt") VALUES
($1, $2, $3, $4, $5, NOW(), NOW()) 
ON CONFLICT (path) DO UPDATE SET contents=$5, "updatedAt"=NOW();`, template.ID, template.IsWorkflow, template.Name, template.Path, template.Contents)
	if err != nil {
		return errors.Wrap(err, "could not add template")
	}
	return nil
}

// Delete deletes a template for an ID
func (s *templatesService) Delete(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.db.Exec(ctx, `DELETE FROM templates WHERE ID=$1;`, ID)
	if err != nil {
		return errors.Wrap(err, "could not delete template")
	}
	return nil
}
