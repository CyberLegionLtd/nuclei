package handlers

import "github.com/labstack/echo/v4"

// GetDashboard swagger:route GET /dashboard dashbaord getDashboard
//
// Returns all the dashboard for the nuclei engine.
func (s *Server) GetDashboard(ctx echo.Context) error {
	return nil
}
