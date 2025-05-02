package database

import (
	"quizapp/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Database Connection Error")
	}
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{}, &models.Score{})
}
