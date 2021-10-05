package db

import (
	"context"

	"github.com/projectdiscovery/nuclei/v2/pkg/web/model"
)

type Templates interface {
	// GetTemplates returns templates for a query
	GetTemplates(ctx context.Context) ([]*model.Template, error)
	// GetTemplateForID returns template for a query ID
	GetTemplateForID(ctx context.Context, templateID string, contents bool) (*model.Template, error)
	// AddTemplate adds a template to the storage. If the template
	// already exists, it is updated.
	AddTemplate(ctx context.Context, template *model.Template) error
	// DeleteTemplateForID deletes a template for an ID
	DeleteTemplateForID(ctx context.Context, templateID string) error
}

type Targets interface {
	// GetTargets returns targets stored for the engine.
	GetTargets(ctx context.Context) ([]*model.Target, error)
	// GetTargetForID returns a specific ID
	GetTargetForID(ctx context.Context, targetID string) (*model.Target, error)
	// AddTarget adds a target to the storage. If the target
	// already exists, it is updated.
	//
	// Append specifies whether the input should be appended or overwritten
	// completely.
	AddTarget(ctx context.Context, target *model.Target, append bool) error
	// DeleteTargetForID deletes a target for an ID
	DeleteTargetForID(ctx context.Context, targetID string) error
}
