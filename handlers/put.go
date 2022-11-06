package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/moaabb/api-postgres/entities"
)

func (m *Handlers) UpdateMovie(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		m.L.Error(err.Error())
		c.JSON(http.StatusBadRequest, Message{
			Message: "Error processing request, check the id",
		})
		return err
	}

	var response entities.MoviePayload

	c.Bind(&response)

	var movie entities.Movie

	movie.ID = id
	movie.Title = response.Title
	movie.Description = response.Description
	movie.Year, _ = strconv.Atoi(response.Year)
	movie.ReleaseDate, _ = time.Parse("2006-01-02", response.ReleaseDate)
	movie.Runtime, _ = strconv.Atoi(response.Runtime)
	movie.Rating, _ = strconv.Atoi(response.Rating)
	movie.MPAARating = response.MPAARating
	movie.UpdatedAt = time.Now()

	err = m.DB.UpdateByID(movie)
	if err != nil {
		m.L.Error(err.Error())
		c.JSON(http.StatusBadRequest, Message{Message: "[ERROR] Couldn't upload movie date, check your request or try again later!"})
		return err
	}

	result, err := m.DB.GetByID(movie.ID)
	if err != nil {
		m.L.Error(err.Error())
		c.JSON(http.StatusBadRequest, Message{Message: "[ERROR] Couldn't upload movie date, check your request or try again later!"})
		return err
	}

	c.JSON(http.StatusOK, result)

	return nil
}
