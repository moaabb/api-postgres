package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moaabb/api-postgres/config"
)

type Message struct {
	Message string `json:"messaage"`
}

type Handlers struct {
	app *config.Application
}

func NewHandlers(a *config.Application) *Handlers {
	return &Handlers{app: a}
}

func (m *Handlers) GetAll(c echo.Context) error {
	c.JSON(http.StatusOK, Message{
		Message: "API is working",
	})

	return nil
}
