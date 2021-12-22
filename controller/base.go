package controller

import (
	"LinkedTo-go/model"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func getAll[P model.Model](c echo.Context, list *[]P) error {
	if err := model.DB.Find(list).Error; err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, list)
}

func getByIndustryType[P model.Model](c echo.Context, list *[]P) error {
	var industryType string
	if err := echo.QueryParamsBinder(c).
		MustString("industryType", &industryType).
		BindError(); err != nil {
		return getAll(c, list)
	}

	if err := model.DB.Where("industry_type = ?", industryType).Find(list).Error; err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, list)

}
