package config

import (
	"os"
	"github.com/joho/godotenv"
)

// Struct chứa tất cả config
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string
}

// Hàm load config từ .env file
func Load() *Config {
	// Load file .env
	godotenv.Load()
	
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "password123"),
		DBName:     getEnv("DB_NAME", "auth_db"),
		ServerPort: getEnv("PORT", "8000"),
		JWTSecret:  getEnv("JWT_SECRET", "my-secret-key"),
	}
}

// Helper function: lấy env variable hoặc default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Tạo database URL từ config
func (c *Config) GetDatabaseURL() string {
	return "postgres://" + c.DBUser + ":" + c.DBPassword + 
		   "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + 
		   "?sslmode=disable"
}