package rest

import (
	"fmt"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/projectdiscovery/nuclei/v2/pkg/web/rest/handlers"
)

// StartApplication starts serving the application on a host-port.
func StartApplication(host string, port int, tls bool, server *handlers.Server) {
	// Echo instance
	e := echo.New()

	scheme := "http"
	if tls {
		scheme = "https"
	}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{fmt.Sprintf("%s://%s:%d", scheme, host, port)},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization"},
	}))
	apiGroup := e.Group("/api/v1")

	// Target management APIs
	apiGroup.GET("/targets", server.GetTargets)          // list all targets
	apiGroup.POST("/targets", server.AddTarget)          // add a new target
	apiGroup.GET("/targets/:id", server.GetTarget)       // get a specific target
	apiGroup.PUT("/targets/:id", server.PutTarget)       // update an existing target list fully
	apiGroup.PATCH("/targets/:id", server.PatchTarget)   // append some input to a target list
	apiGroup.DELETE("/targets/:id", server.DeleteTarget) // delete a target from the list

	// Template management APIs
	apiGroup.GET("/templates", server.GetTemplates)                 // list all templates (optional workflows parameter)
	apiGroup.POST("/templates", server.AddTemplate)                 // add a new template
	apiGroup.GET("/templates/:id", server.GetTemplateForID)         // get a specific template with id
	apiGroup.PUT("/templates/:id", server.UpdateTemplateForID)      // update an existing template with id
	apiGroup.POST("/templates/:id/execute", server.ExecuteTemplate) // execute a specific template with id (optional debug)

	// Scan management APIs
	apiGroup.GET("/scans", server.GetScans)          // list all scans
	apiGroup.POST("/scans", server.AddScan)          // add a new scan
	apiGroup.GET("/scans/:id", server.GetScan)       // get a specific scan
	apiGroup.PUT("/scans/:id", server.PutScan)       // update an existing scan status (pause,stop)
	apiGroup.DELETE("/scans/:id", server.DeleteScan) // delete a scan from the list

	// Setting management APIs
	apiGroup.GET("/settings", server.GetSettings) // get all settings

	// Dashboard APIs
	apiGroup.GET("/dashboard", server.GetDashboard) // get all dashboard

	// Misc APIs
	apiGroup.GET("/usage", server.GetUsage) // gets the usage metrics

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}
