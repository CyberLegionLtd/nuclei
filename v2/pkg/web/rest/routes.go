package rest

import (
	"fmt"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func StartApplication(host string, port int, tls bool) {
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
	apiGroup.GET("/targets", handlerFunc)         // list all targets
	apiGroup.POST("/targets", handlerFunc)        // add a new target
	apiGroup.GET("/targets/:id", handlerFunc)     // get a specific target
	apiGroup.PUT("/targets/:id", handlerFunc)     // update an existing target list fully
	apiGroup.PATCH("/targets/:id", handlerFunc)   // append some input to a target list
	apiGroup.DELETE("/targets/:id", handlerFunc)  // delete a target from the list
	apiGroup.GET("/targets/:id/raw", handlerFunc) // get a target list in raw format

	// Scan management APIs
	apiGroup.GET("/scans", handlerFunc)                 // list all scans
	apiGroup.POST("/scans", handlerFunc)                // add a new scan
	apiGroup.GET("/scans/:id", handlerFunc)             // get a specific scan
	apiGroup.PUT("/scans/:id", handlerFunc)             // update an existing scan status (pause,stop)
	apiGroup.DELETE("/scans/:id", handlerFunc)          // delete a scan from the list
	apiGroup.GET("/scans/:id/errors", handlerFunc)      // get a list of errors for a scan
	apiGroup.GET("/scans/:id/matches", handlerFunc)     // get a list of matches for a scan
	apiGroup.GET("/scans/:id/results/:id", handlerFunc) // get a specific result id for a scan id

	// Template management APIs
	apiGroup.GET("/templates", handlerFunc)               // list all templates
	apiGroup.GET("/workflows", handlerFunc)               // list all workflows
	apiGroup.POST("/templates", handlerFunc)              // add a new template
	apiGroup.POST("/workflows", handlerFunc)              // add a new workflow
	apiGroup.GET("/templates/:path", handlerFunc)         // get a specific template with path
	apiGroup.GET("/workflows/:path", handlerFunc)         // get a specific workflow with path
	apiGroup.PUT("/templates/:path", handlerFunc)         // update an existing template with path
	apiGroup.PUT("/workflows/:path", handlerFunc)         // update an existing workflow with path
	apiGroup.GET("/templates/:path/execute", handlerFunc) // execute a specific template with path (optional debug)
	apiGroup.GET("/workflows/:path/execute", handlerFunc) // execute a specific workflow with path (optional debug)

	// Setting management APIs
	apiGroup.GET("/settings", handlerFunc) // get all settings

	// Dashboard APIs
	apiGroup.GET("/dashboard", handlerFunc) // get all dashboard

	// Misc APIs
	apiGroup.GET("/usage", handlerFunc) // gets the usage metrics

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}

func handlerFunc(echo.Context) error {
	return nil
}
