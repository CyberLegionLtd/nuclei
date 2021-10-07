package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
)

type Templates interface {
	// GetTemplates returns all templates in the database
	List(callback func(template *model.Template)) error
	// Get returns the contents for a template ID
	Get(ID string) (string, error)
	// Delete deletes a template for an ID
	Delete(ID string) error
	// Add adds a template to the db.
	// If the template already exists for a path, it's contents are updated.
	Add(template *model.Template) error
}

type Targets interface {
	// List returns all targets in db
	List(callback func(template *model.Target)) error
	// Delete deletes a target for an ID
	Delete(ID string) error
	// GetTargetForID returns a specific ID
	GetTargetForID(ctx context.Context, targetID string) (*model.Target, error)
	// AddTarget adds a target to the storage. If the target
	// already exists, it is updated.
	//
	// Append specifies whether the input should be appended or overwritten
	// completely.
	AddTarget(ctx context.Context, target *model.Target, append bool) error
}

type DB struct {
	db *pgxpool.Pool
}

func New(connStr string) (*DB, error) {
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return &DB{db: conn}, nil
}

func (d *DB) DB() *pgxpool.Pool {
	return d.db
}

func (d *DB) Close() {
	d.db.Close()
}
