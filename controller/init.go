package controller

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
}

func StartServer() {
	e.Logger.Fatal(e.Start(":8080"))
}
