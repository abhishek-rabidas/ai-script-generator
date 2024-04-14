package views

import "github.com/labstack/echo/v4"

type CommonApiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GenerateApiResponse(ctx echo.Context, status int, message string, data any) error {
	return ctx.JSON(status, CommonApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
