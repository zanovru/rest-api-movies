package db

import (
	"github.com/zanovru/rest-api-movies/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL: %v", err)
	}
	err = db.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatalf("Failed to automigrate database: %v", err)
	}
	return db
}
