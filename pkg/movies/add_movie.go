package movies

import (
	"github.com/gin-gonic/gin"
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"net/http"
)

type AddMovieRequestBody struct {
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Director    string `json:"director"`
	Description string `json:"description"`
}

func (h handler) AddMovie(ctx *gin.Context) {
	body := AddMovieRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var movie models.Movie

	movie.Title = body.Title
	movie.Year = body.Year
	movie.Director = body.Director
	movie.Description = body.Description

	if result := h.DB.Create(&movie); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

}
