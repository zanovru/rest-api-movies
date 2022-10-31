package movies

import (
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movie

	if result := h.DB.First(&movie, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&movie)

	c.Status(http.StatusOK)
}
