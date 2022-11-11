package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zanovru/rest-api-movies/pkg/common/db"
	"github.com/zanovru/rest-api-movies/pkg/movies"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	r := gin.Default()
	dbORM := db.Init()

	movies.RegisterRoutes(r, dbORM)

	r.Run(":3000")
}
