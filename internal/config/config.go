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

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return
		}

		config = &Config{
			ServerAddress: os.Getenv("SERVER_ADDRESS"),
			DBHost:        os.Getenv("DB_HOST"),
			DBPort:        os.Getenv("DB_PORT"),
			DBUser:        os.Getenv("DB_USER"),
			DBPassword:    os.Getenv("DB_PASSWORD"),
			DBName:        os.Getenv("DB_NAME"),
		}
	})

	return config, err
}
