package main

import (
	"ai-script-generator/config"
	"ai-script-generator/handler"
	"ai-script-generator/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	scriptService := service.NewScriptService()
	handler.NewScriptHandler(apiGroup.Group("/script"), scriptService)

	config.SetupDatabase()

	err := e.Start(":5000")
	if err != nil {
		panic(err)
	}
}
