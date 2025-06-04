package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	RedisAddr string
	RedisPass string
	JwtSecret string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env file not found")
	}

	cfg := &Config{
		Port:      os.Getenv("PORT"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		RedisPass: os.Getenv("REDIS_PASS"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT not set")
	}

	return cfg, nil
}
