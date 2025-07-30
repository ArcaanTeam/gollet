package config

import (
	"fmt"
	"gollet/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}

	log.Println("Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database.", err)
	}
	log.Println("Database migration completed")
}
