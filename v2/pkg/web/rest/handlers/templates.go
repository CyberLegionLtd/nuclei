package handlers

import "github.com/labstack/echo/v4"

// GetTemplates swagger:route GET /templates templates getTemplates
//
// Returns all the templates available to the engine.
func (s *Server) GetTemplates(ctx echo.Context) error {
	
	return nil
}

// GetTemplateForID swagger:route GET /templates/:id templates getTemplateForID
//
// Returns a template for a ID.
func (s *Server) GetTemplateForID(ctx echo.Context) error {
	return nil
}

// AddTemplate swagger:route POST /templates templates addTemplate
//
// Returns all the templates available to the engine.
func (s *Server) AddTemplate(ctx echo.Context) error {
	return nil
}

// UpdateTemplateForID swagger:route PUT /templates/:id templates updateTemplateForID
//
// Update a template for a path.
func (s *Server) UpdateTemplateForID(ctx echo.Context) error {
	return nil
}

// ExecuteTemplate swagger:route POST /templates/:id/execute templates executeTemplate
//
// Executes a template.
func (s *Server) ExecuteTemplate(ctx echo.Context) error {
	return nil
}
