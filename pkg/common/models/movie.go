package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	Year        int    `json:"year"`
	Director    string `json:"director"`
	Description string `json:"description"`
}
