package handler

import (
	"ai-script-generator/service"
	"ai-script-generator/views"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(e *echo.Group, service *service.UserService) *UserHandler {
	handler := UserHandler{service: service}

	e.POST("/", handler.createAccount)
	e.POST("/voice", handler.captureVoice)

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

func (u UserHandler) captureVoice(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return views.GenerateApiResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	ext := strings.Split(file.Filename, ".")[1]

	if ext != "mp3" {
		return views.GenerateApiResponse(c, http.StatusInternalServerError, "unsupported Audio Format only .mp3 is supported", nil)
	}

	src, err := file.Open()
	if err != nil {
		log.Error(err)
		return views.GenerateApiResponse(c, http.StatusInternalServerError, "failed to store audio", nil)
	}

	defer src.Close()

	dst, err := os.Create("data/" + file.Filename)
	if err != nil {
		log.Error(err)
		return views.GenerateApiResponse(c, http.StatusInternalServerError, "failed to store audio", nil)
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Error(err)
		return views.GenerateApiResponse(c, http.StatusInternalServerError, "failed to store audio", nil)
	}

	err = u.service.AnalyzeVoice("data/"+file.Filename, c.FormValue("uid"))

	if err != nil {
		return views.GenerateApiResponse(c, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return views.GenerateApiResponse(c, http.StatusOK, "Audio Analysis Successful", nil)
	}
}
