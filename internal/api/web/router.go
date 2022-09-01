package web

import (
	"login-app/internal/api/web/user/handler"

	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	handler := handler.NewController()

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	g := e.Group("/v1")

	{
		g.POST("/users", handler.Post)
	}

	return e
}
