package user

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewController() *Handler {
	return &Handler{}
}

func (h *Handler) Post(c echo.Context) error {
	return nil
}
