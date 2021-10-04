package handlers

import "github.com/labstack/echo"

// GetScans swagger:route GET /scans scans getScans
//
// Returns all the scans for the engine.
func (s *Server) GetScans(ctx echo.Context) error {
	return nil
}

// AddScan swagger:route POST /scans targets addScan
//
// Adds a new scan to the scan queue.
func (s *Server) AddScan(ctx echo.Context) error {
	return nil
}

// GetScan swagger:route GET /scans/:id targets getScan
//
// Returns a scan for an ID.
func (s *Server) GetScan(ctx echo.Context) error {
	return nil
}

// PutScan swagger:route PUT /scans/:id targets putScan
//
// Updates a scan configuration.
func (s *Server) PutScan(ctx echo.Context) error {
	return nil
}

// DeleteScan swagger:route DELETE /scans/:id targets deleteScan
//
// Deletes a scan from the list.
func (s *Server) DeleteScan(ctx echo.Context) error {
	return nil
}

// GetScanErrors swagger:route GET /scans/:id/errors targets getScanErrors
//
// Returns errors for a scan ID.
func (s *Server) GetScanErrors(ctx echo.Context) error {
	return nil
}

// GetScanMatches swagger:route GET /scans/:id/matches targets getScanMatches
//
// Returns matches for a scan ID.
func (s *Server) GetScanMatches(ctx echo.Context) error {
	return nil
}

// GetScanResultForID swagger:route GET /scans/:id/results/:id targets getScanResultForID
//
// Returns scan results for an ID.
func (s *Server) GetScanResultForID(ctx echo.Context) error {
	return nil
}
