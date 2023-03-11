package main

import (
	"meli/config"
	"meli/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.HealthCheck)

	config.Connect()

	e.POST("/data", handlers.CreateData)
	e.GET("/data", handlers.GellAllData)

	e.Logger.Fatal(e.Start(":8080"))

}
