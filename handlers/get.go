package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Handlers) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		m.L.Error(err.Error())
		c.JSON(http.StatusBadRequest, Message{
			Message: "Error processing request, check the id",
		})
		return err
	}

	movie, err := m.DB.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message{
			Message: err.Error(),
		})
		m.L.Error(err.Error())
		return err
	}

	c.JSON(http.StatusOK, movie)

	return nil
}

func (m *Handlers) GetAll(c echo.Context) error {
	movies, err := m.DB.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, Message{"Could not fetch the data, try again later!"})
		m.L.Trace(err.Error())
		return err
	}

	c.JSON(http.StatusOK, movies)

	return nil
}
