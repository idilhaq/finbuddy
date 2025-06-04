package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	// Only load .env if not running inside Docker
	// (you can also check for a special env var in docker-compose)
	_, inDocker := os.LookupEnv("DOCKER_ENV")
	if !inDocker {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️  No .env file found, continuing without it")
		} else {
			log.Println("✅ Loaded .env file")
		}
	} else {
		log.Println("✅ Running inside Docker, skip loading .env")
	}

	sslmode := "require"
	if os.Getenv("APP_ENV") == "development" || inDocker {
		sslmode = "disable"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	DB = db
	return nil
}
