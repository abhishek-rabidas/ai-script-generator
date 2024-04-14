package handler

import (
	"ai-script-generator/service"
	"ai-script-generator/views"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ScriptHandler struct {
	service *service.ScriptService
}

func NewScriptHandler(e *echo.Group, service *service.ScriptService) *ScriptHandler {
	handler := ScriptHandler{service: service}

	e.POST("/create", handler.createScript)

	return &handler

}

func (s ScriptHandler) createScript(c echo.Context) error {
	req := views.NewScriptRequest{}

	err := c.Bind(&req)

	if err != nil {
		return views.GenerateApiResponse(c, http.StatusBadRequest, "check payload", nil)
	}

	res, err := s.service.GenerateNewScript(req)

	if err != nil {
		return views.GenerateApiResponse(c, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return views.GenerateApiResponse(c, http.StatusOK, "Script created", res)
	}
}
