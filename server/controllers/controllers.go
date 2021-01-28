package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Ping ...
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong...")
}
