package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppPort   string
	DBURL     string
	RedisAddr string
	RedisPass string
	RedisDB   int
}

func LoadConfig() *Config {
	redisDBStr := getEnv("REDIS_DB", "0")
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Println("Invalid REDIS_DB, defaulting to 0")
		redisDB = 0
	}

	cfg := &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		DBURL:     getEnv("DB_URL", "postgres://postgres:S0@130105@db:5432/url_db?sslmode=disable"),
		RedisAddr: getEnv("REDIS_ADDR", "redis:6379"),
		RedisPass: getEnv("REDIS_PASSWORD", ""),
		RedisDB:   redisDB,
	}

	validate(cfg)

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func validate(cfg *Config) {
	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}
	if cfg.RedisAddr == "" {
		log.Fatal("REDIS_ADDR is required")
	}
}
