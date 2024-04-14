package handler

import (
	"ai-script-generator/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(e *echo.Group, service *service.UserService) *UserHandler {
	handler := UserHandler{service: service}

	e.POST("/", handler.createAccount)

	return &handler
}

func (u UserHandler) createAccount(c echo.Context) error {
	return c.String(200, "Account creation")
}
