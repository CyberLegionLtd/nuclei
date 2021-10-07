package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
)

type targetsService struct {
	db *pgxpool.Pool
}

func newTargetsService(db *pgxpool.Pool) *targetsService {
	return &targetsService{db: db}
}

// List returns all targets in db
func (s *targetsService) List(callback func(target *model.Target)) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, `SELECT
id,
name,
"rawPath",
"totalHosts",
"createdAt",
"updatedAt" FROM targets;`)
	if err != nil {
		return errors.Wrap(err, "could not list targets")
	}
	defer rows.Close()

	target := &model.Target{}
	for rows.Next() {
		if scanErr := rows.Scan(&target.ID, &target.Name, &target.RawPath, &target.TotalHosts, &target.CreatedAt, &target.UpdatedAt); scanErr != nil {
			err = scanErr
		}
		callback(target)
	}
	return err
}

// Add adds a target to the db.
func (s *targetsService) Add(target *model.Target) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.db.Exec(ctx, `INSERT INTO targets (
id,
name,
"rawPath",
"totalHosts",
"createdAt",
"updatedAt") VALUES
($1, $2, $3, $4, NOW(), NOW()) 
ON CONFLICT ("rawPath") DO UPDATE SET "updatedAt"=NOW();`, target.ID, target.Name, target.RawPath, target.TotalHosts)
	if err != nil {
		return errors.Wrap(err, "could not add template")
	}
	return nil
}

// Delete deletes a target for an ID
func (s *targetsService) Delete(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.db.Exec(ctx, `DELETE FROM targets WHERE ID=$1;`, ID)
	if err != nil {
		return errors.Wrap(err, "could not delete target")
	}
	return nil
}
