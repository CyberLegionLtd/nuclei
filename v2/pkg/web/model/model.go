package model

import (
	"time"
)

// Template is a template stored for nuclei engine.
type Template struct {
	// ID is the ID of the template
	ID string `json:"id,omitempty"`
	// Path is the path of the stored template
	Path string `json:"path,omitempty"`
	// Name is the name of the specified template
	Name string `json:"name,omitempty"`
	// IsWorkflow specifies whether the template is a workflow.
	IsWorkflow bool `json:"isWorkflow,omitempty"`
	// CreatedAt is the time the target was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt is the time the target was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Contents is the contents of the template file.
	Contents string `json:"contents,omitempty"`
}

// Target is a target list that can be scanned with templates.
type Target struct {
	// ID is the ID of the target
	ID string `json:"id,omitempty"`
	// Name is the name of the target
	Name string `json:"name,omitempty"`
	// RawPath is the path to the target file.
	RawPath string `json:"rawPath,omitempty"`
	// TotalHosts count available in the target list
	TotalHosts int64 `json:"totalHosts,omitempty"`
	// CreatedAt is the time the target was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt is the time the target was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
