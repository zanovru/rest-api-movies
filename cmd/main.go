package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zanovru/rest-api-movies/pkg/common/db"
	"github.com/zanovru/rest-api-movies/pkg/movies"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbURL := viper.Get("DB_URL").(string)

	r := gin.Default()
	dbORM := db.Init(dbURL)

	movies.RegisterRoutes(r, dbORM)

	r.Run(port)
}
