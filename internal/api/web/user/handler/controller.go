package handler

import (
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
		return c.JSON(400, err)
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return c.JSON(400, err)
	}

	id, err := h.service.Create(request)

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, id)
}
