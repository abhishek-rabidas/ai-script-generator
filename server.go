package main

import (
	"ai-script-generator/config"
	"ai-script-generator/handler"
	"ai-script-generator/service"
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

	userService := service.NewUserService()
	handler.NewUserHandler(e.Group("/user"), userService)

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
