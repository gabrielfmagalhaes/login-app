package handler

import (
	"login-app/internal/api/common"
	"login-app/internal/api/web/user/dto"
	"login-app/internal/core/services"
	"login-app/platform/validator"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *services.Service
}

func NewController(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(c echo.Context) error {
	request := new(dto.CreateUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(common.NewBadRequestResponse(err))
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return c.JSON(common.NewBadRequestResponse(err))
	}

	id, err := h.service.Create(request)

	if err != nil {
		return c.JSON(common.NewBadRequestResponse(err))
	}

	return c.JSON(common.NewCreatedResponse(dto.CreateUserResponse{Id: id}))
}
