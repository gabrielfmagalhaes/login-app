package user

import (
	"login-app/api/common"
	"login-app/api/user/request"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewController() *Handler {
	return &Handler{}
}

func (h *Handler) Post(c echo.Context) error {
	request := new(request.CreateUserRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	return c.JSON(common.NewCreatedResponse())
}
