package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress  string
	LoaderFilePath string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
}

var (
	config *Config
	once   sync.Once
)

// LoadConfig carga las variables de entorno solo una vez
func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if err = godotenv.Load("/Users/judiazgutier/Documents/GoWeb/Sprints/dev/wave15-grupo4/.env"); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return
		}

		config = &Config{
			ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
			DBHost:        getEnv("DB_HOST", "localhost"),
			DBPort:        getEnv("DB_PORT", "3306"),
			DBUser:        getEnv("DB_USER", "root"),
			DBPassword:    getEnv("DB_PASSWORD", "password"),
			DBName:        getEnv("DB_NAME", "grupo4"),
		}
	})

	return config, err
}

// getEnv obtiene una variable de entorno o un valor por defecto
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
