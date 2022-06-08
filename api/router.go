package api

import (
	"login-app/api/user"

	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	handler := user.NewController()

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	g := e.Group("/v1")

	{
		g.POST("/users", handler.Post)
	}

	return e
}
