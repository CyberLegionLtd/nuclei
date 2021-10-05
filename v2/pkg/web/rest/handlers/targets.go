package handlers

import "github.com/labstack/echo/v4"

// GetTargets swagger:route GET /targets targets getTargets
//
// Returns all the targets for the target storage.
func (s *Server) GetTargets(ctx echo.Context) error {
	return nil
}

// AddTarget swagger:route POST /targets targets addTarget
//
// Adds a new target to the target storage.
func (s *Server) AddTarget(ctx echo.Context) error {
	return nil
}

// GetTarget swagger:route GET /targets/:id targets getTarget
//
// Returns a target for an ID.
func (s *Server) GetTarget(ctx echo.Context) error {
	// Optional parameter - raw returning raw data for target.
	return nil
}

// PutTarget swagger:route PUT /targets/:id targets putTarget
//
// Updates a target list with a new list.
func (s *Server) PutTarget(ctx echo.Context) error {
	return nil
}

// PatchTarget swagger:route PATCH /targets/:id targets patchTarget
//
// Appends some input to a target list.
func (s *Server) PatchTarget(ctx echo.Context) error {
	return nil
}

// DeleteTarget swagger:route DELETE /targets/:id targets deleteTarget
//
// Removes a target for an ID.
func (s *Server) DeleteTarget(ctx echo.Context) error {
	return nil
}
