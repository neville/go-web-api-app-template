package server

import (
	"github.com/labstack/echo/v4"
	"github.com/neville/go-web-api-app-template/server/routes"
)

// Start ...
func Start() {
	e := echo.New()

	routes.Register(e)

	e.Start(":8080")
}
