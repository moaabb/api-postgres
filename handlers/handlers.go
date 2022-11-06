package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/moaabb/api-postgres/config"
	"github.com/moaabb/api-postgres/driver"
)

type Message struct {
	Message string `json:"messaage"`
}

type Handlers struct {
	a  *config.Application
	DB driver.DBModel
}

func NewHandlers(a *config.Application) *Handlers {
	return &Handlers{a, driver.NewDB(a.DBModel)}
}

func (m *Handlers) GetAll(c echo.Context) error {
	movies, err := m.DB.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, Message{"Could not fetch the data, try again later!"})
		m.a.L.Error(err.Error())
		return err
	}

	c.JSON(http.StatusOK, movies)

	return nil
}
