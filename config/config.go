package config

import (
	"fmt"
	"os"
)

type Config struct {
	PostgresURL string
	RedisAddr   string
}

func Load() (*Config, error) {
	pgURL := os.Getenv("POSTGRES_URL")
	if pgURL == "" {
		return nil, fmt.Errorf("POSTGRES_URL not set")
	}
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	return &Config{
		PostgresURL: pgURL,
		RedisAddr:   redisAddr,
	}, nil
}
