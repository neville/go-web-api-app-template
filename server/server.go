package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neville/go-web-api-app-template/server/routes"
)

// Start ...
func Start() {
	// Create instance
	e := echo.New()

	// Register middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register routes
	routes.Register(e)

	// Start HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}
