package config

import (
	"os"
)

type Config struct {
	AppPort  string
	DBDSN    string
	LogLevel string
}

func Load() Config {
	return Config{
		AppPort:  getEnv("APP_PORT", "8080"),
		DBDSN:    getEnv("DB_DSN", ""),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func (c Config) DBDsnEmpty() bool { return c.DBDSN == "" }
