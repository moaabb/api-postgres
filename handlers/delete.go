package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Handlers) DeleteMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		m.L.Error(err.Error())
		c.JSON(http.StatusBadRequest, Message{
			Message: "Error processing request, check the id",
		})
		return err
	}

	err = m.DB.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message{
			Message: err.Error(),
		})
		m.L.Error(err.Error())
		return err
	}

	c.JSON(http.StatusOK, Message{"Resource deleted!"})

	return nil
}
