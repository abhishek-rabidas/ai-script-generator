package main

import (
	"ai-script-generator/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func SetupServer() {
	e := echo.New()
	e.Use(middleware.CORS())

	apiGroup := e.Group("/api")

	apiGroup.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up")
	})

	config.SetupDatabase()

	err := e.Start(":5000")
	if err != nil {
		panic(err)
	}
}
