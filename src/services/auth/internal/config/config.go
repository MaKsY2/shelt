package config

import "os"

type Config struct {
	ServerPort string
	DB         string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		ServerPort: getEnv("AUTH_SERVICE_PORT", "8000"),
		DB: "postgres://" +
			getEnv("DB_USER", "postgres") + ":" +
			getEnv("DB_PASSWORD", "password") + "@" +
			getEnv("DB_HOST", "localhost") + ":" +
			getEnv("DB_PORT", "5432") + "/" +
			getEnv("DB_NAME", "mydb") + "?sslmode=disable",
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
