package config

import "os"

type Config struct {
	PostgresURL string
	RedisAddr   string
	JWTSecret   string
}

func Load() (*Config, error) {
	return &Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
