package handlers

import "github.com/labstack/echo"

// GetTemplates swagger:route GET /templates templates getTemplates
//
// Returns all the templates available to the engine.
func (s *Server) GetTemplates(ctx echo.Context) error {
	return nil
}

// GetWorkflows swagger:route GET /workflows templates getWorkflows
//
// Returns all the workflows available to the engine.
func (s *Server) GetWorkflows(ctx echo.Context) error {
	return nil
}

// GetTemplates swagger:route GET /templates/:path templates getTemplateForPath
//
// Returns a template for a path.
func (s *Server) GetTemplateForID(ctx echo.Context) error {
	return nil
}

// GetWorkflowForID swagger:route GET /workflows/:path templates getWorkflowForPath
//
// Returns a workflow for a path.
func (s *Server) GetWorkflowForID(ctx echo.Context) error {
	return nil
}

// AddTemplate swagger:route POST /templates templates addTemplate
//
// Returns all the templates available to the engine.
func (s *Server) AddTemplate(ctx echo.Context) error {
	return nil
}

// AddWorkflow swagger:route POST /workflows templates addWorkflow
//
// Returns all the workflows available to the engine.
func (s *Server) AddWorkflow(ctx echo.Context) error {
	return nil
}

// UpdateTemplateForID swagger:route PUT /templates/:path templates updateTemplateForPath
//
// Update a template for a path.
func (s *Server) UpdateTemplateForID(ctx echo.Context) error {
	return nil
}

// UpdateWorkflowForID swagger:route PUT /workflows/:path templates updateWorkflowForPath
//
// Update a workflow for a path.
func (s *Server) UpdateWorkflowForID(ctx echo.Context) error {
	return nil
}

// ExecuteTemplate swagger:route GET /templates/:path/execute templates executeTemplate
//
// Executes a template.
func (s *Server) ExecuteTemplate(ctx echo.Context) error {
	return nil
}

// ExecuteWorkflow swagger:route GET /workflows/:path/execute templates executeWorkflow
//
// Executes a workflow.
func (s *Server) ExecuteWorkflow(ctx echo.Context) error {
	return nil
}
