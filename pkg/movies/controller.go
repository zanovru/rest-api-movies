package movies

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/movies")
	routes.POST("/", h.AddMovie)
	routes.GET("/", h.GetMovies)
	routes.GET("/:id", h.GetMovie)
	routes.PUT("/:id", h.UpdateMovie)
	routes.DELETE("/:id", h.DeleteMovie)
}
