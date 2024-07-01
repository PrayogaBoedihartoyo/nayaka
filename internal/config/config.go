package config

import (
    "github.com/joho/godotenv"
    "fmt"
    "os"
)

type Config struct {
    DatabaseURL string
}

func Load() (*Config, error) {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("Error loading .env file: %v", err)
    }

    // Read values from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Construct database URL
    dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", 
        dbUser, dbPassword, dbHost, dbPort, dbName)

    return &Config{
        DatabaseURL: dbURL,
    }, nil
}