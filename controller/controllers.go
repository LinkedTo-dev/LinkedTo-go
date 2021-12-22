package controller

import (
	"LinkedTo-go/model"
	"github.com/labstack/echo/v4"
)

func getMaps(c echo.Context) error {
	return getAll(c, &[]model.Map{})
}

func getNews(c echo.Context) error {
	return getByIndustryType(c, &[]model.News{})
}

func getPolicies(c echo.Context) error {
	return getAll(c, &[]model.Policy{})
}

func getSpecialists(c echo.Context) error {
	return getByIndustryType(c, &[]model.Specialist{})
}

func getMapStatistics(c echo.Context) error {
	return getByIndustryType(c, &[]model.MapStatistic{})
}

func getDataStatistics(c echo.Context) error {
	return getByIndustryType(c, &[]model.DataStatistic{})
}
