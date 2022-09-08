package web

import (
	"login-app/internal/api/web/user/handler"
	"login-app/internal/core/services"

	"github.com/labstack/echo/v4"
)

func SetupRouter(userService *services.Service) *echo.Echo {
	e := echo.New()

	handler := handler.NewController(userService)

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	g := e.Group("/v1")

	{
		g.POST("/users", handler.Post)
	}

	return e
}
