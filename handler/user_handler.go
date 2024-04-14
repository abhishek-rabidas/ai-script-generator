package handler

import (
	"ai-script-generator/service"
	"ai-script-generator/views"
	"github.com/labstack/echo/v4"
	"net/http"
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
	err := c.Bind(&req)

	if err != nil {
		return views.GenerateApiResponse(c, http.StatusBadRequest, "check payload", nil)
	}

	res, err := u.service.CreateUser(req)

	if err != nil {
		return views.GenerateApiResponse(c, http.StatusInternalServerError, "failed to create user", nil)
	} else {
		return views.GenerateApiResponse(c, http.StatusOK, "User created", res)
	}

}
