package config

import "os"

type Config struct {
	Port  string
	DBDsn string
}

func Load() *Config {
	return &Config{
		Port:  getEnv("PORT", "8081"),
		DBDsn: getEnv("DB_DSN", "appuser:apppass@tcp(localhost:3306)/achievement_leveling?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
