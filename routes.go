package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/moaabb/api-postgres/handlers"
)

func routes() *echo.Echo {
	e := echo.New()

	h := handlers.NewHandlers(&app)

	// e.Use(middleware.Logger())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			app.L.Info(fmt.Sprintf("%s %s", c.Request().Method, c.Request().URL.Path))
			return next(c)
		}
	})

	e.GET("/api", h.GetAll)

	return e
}
