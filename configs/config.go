package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config структура конфигурации приложения
type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

// DbConfig структура для БД
type DbConfig struct {
	Dsn string
}

// AuthConfig структура для авторизации
type AuthConfig struct {
	Secret string
}

// LoadConfig загрузка конфигурации
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
