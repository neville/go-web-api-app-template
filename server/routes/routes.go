package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/neville/go-web-api-app-template/server/controllers"
)

// Register ...
func Register(e *echo.Echo) {
	e.GET("/", controllers.Ping)
}
