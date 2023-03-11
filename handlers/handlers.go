package handlers

import (
	"meli/controller"
	"meli/model"

	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Status Ok!!")
}

func GellAllData(c echo.Context) error {
	data, err := controller.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func CreateData(c echo.Context) error {
	data := model.Data{}
	c.Bind(&data)
	err := controller.InsertData(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
