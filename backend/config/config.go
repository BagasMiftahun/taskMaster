package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfigStruct struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
	Port string
}

var AppConfig AppConfigStruct

// Load menginisialisasi pengaturan aplikasi dari file .env.
func Load() {
	// Memuat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing with defaults")
	}

	// Mengisi AppConfig dari variabel lingkungan
	AppConfig.Database.User = os.Getenv("DB_USER")
	AppConfig.Database.Password = os.Getenv("DB_PASSWORD")
	AppConfig.Database.Host = os.Getenv("DB_HOST")
	AppConfig.Database.Port = os.Getenv("DB_PORT")
	AppConfig.Database.Name = os.Getenv("DB_NAME")
	AppConfig.Port = os.Getenv("APP_PORT")

	// Validasi pengaturan yang diperlukan
	if AppConfig.Database.User == "" || AppConfig.Database.Password == "" {
		log.Fatal("Database user and password must be set.")
	}
}
