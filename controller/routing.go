package controller

import (
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/map", getMaps)
	api.GET("/policy", getPolicies)
	api.GET("/specialist", getSpecialists)
	api.GET("/news", getNews)
	api.GET("/mapStatistic", getMapStatistics)
	api.GET("/dataStatistic", getDataStatistics)
}
