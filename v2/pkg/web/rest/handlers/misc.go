package handlers

import "github.com/labstack/echo/v4"

// GetUsage swagger:route GET /usage misc getUsage
//
// Returns the usage statistics for the web API.
func (s *Server) GetUsage(ctx echo.Context) error {
	return nil
}
