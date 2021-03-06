package handler

import (
	"login-app/internal/api/common"
	"login-app/internal/api/user/dto"
	"login-app/platform/validator"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewController() *Handler {
	return &Handler{}
}

func (h *Handler) Post(c echo.Context) error {
	request := new(dto.CreateUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(common.NewBadRequestResponse(err))
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return c.JSON(common.NewBadRequestResponse(err))
	}

	return c.JSON(common.NewCreatedResponse(request.Email))
}
