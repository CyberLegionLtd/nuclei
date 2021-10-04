package models

import "github.com/go-openapi/strfmt"

// Target is a target list that can be scanned with templates.
// swagger:model target
type Target struct {
	// ID is the ID of the target
	ID strfmt.UUID4 `json:"id,omitempty"`
	// Name is the name of the target
	Name string `json:"name,omitempty"`
	// LastUpdated time of the input list.
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`
	// TotalHosts count available in the target list
	TotalHosts int64 `json:"totalHosts,omitempty"`
}
