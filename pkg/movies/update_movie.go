package movies

import (
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateMovieRequestBody struct {
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Director    string `json:"director"`
	Description string `json:"description"`
}

func (h handler) UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMovieRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var movie models.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	movie.Title = body.Title
	movie.Year = body.Year
	movie.Director = body.Director
	movie.Description = body.Description

	h.DB.Save(&movie)

	c.JSON(http.StatusOK, &movie)
}
