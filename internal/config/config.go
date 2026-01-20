package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	HTTPAddr string
	Version  string
}

func Load() Config {
	_ = godotenv.Load() // in prod se manca va bene

	return Config{
		Env:      getEnv("APP_ENV", "dev"),
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
		Version:  getEnv("APP_VERSION", "0.1.0"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
