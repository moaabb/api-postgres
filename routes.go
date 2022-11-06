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

	e.GET("/api/movies", h.GetAll)
	e.GET("/api/movies/:id", h.GetByID)

	e.POST("/api/movies", h.InsertMovie)

	e.PUT("/api/movies/:id", h.UpdateMovie)

	e.DELETE("/api/movies/:id", h.DeleteMovie)

	return e
}
