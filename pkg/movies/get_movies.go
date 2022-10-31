package movies

import (
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetMovies(c *gin.Context) {
	var movies []models.Movie

	if result := h.DB.Find(&movies); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &movies)
}
