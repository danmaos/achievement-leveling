package config

import "os"

type Config struct {
	Port               string
	JWTSecret          string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	AchievementSvcURL  string
	FrontendURL        string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("PORT", "8080"),
		JWTSecret:          getEnv("JWT_SECRET", "dev-secret-change-me"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		AchievementSvcURL:  getEnv("ACHIEVEMENT_SVC_URL", "http://localhost:8081"),
		FrontendURL:        getEnv("FRONTEND_URL", "http://localhost:3000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
