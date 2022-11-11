package db

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {

	user := viper.Get("DB_USER").(string)
	pass := viper.Get("DB_PASSWORD").(string)
	name := viper.Get("DB_NAME").(string)

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@db:5432/%s", user, pass, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL: %v", err)
	}
	err = db.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatalf("Failed to automigrate database: %v", err)
	}
	return db
}
