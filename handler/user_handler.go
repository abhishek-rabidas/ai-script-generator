package handler

import (
	"ai-script-generator/service"
	"ai-script-generator/views"
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
	req := views.UserCreateRequest{}
	c.Bind(&req)
	return c.JSON(200, u.service.CreateUser(req))
}
