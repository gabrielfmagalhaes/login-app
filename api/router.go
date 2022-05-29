package api

import (
	"login-app/api/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, handler *user.Handler) {

	e.POST("/v1/users", handler.Post)
}
