package handlers

import "github.com/labstack/echo/v4"

// GetSettings swagger:route GET /settings settings getSettings
//
// Returns all the settings for the nuclei engine.
func (s *Server) GetSettings(ctx echo.Context) error {
	return nil
}
