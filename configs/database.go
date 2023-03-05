package configs

import (
	"go-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=123456789 dbname=pustaka_api port=5432 sslmode=disable"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connected")

}