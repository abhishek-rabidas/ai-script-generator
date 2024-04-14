package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetupServer() {
	e := echo.New()
	e.Use(middleware.CORS())

	apiGroup := e.Group("/api")

	apiGroup.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up")
	})

	err := e.Start(":5000")
	if err != nil {
		panic(err)
	}
}
