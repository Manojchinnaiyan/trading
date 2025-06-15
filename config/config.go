package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseURL          string
	RedisURL             string
	RedisPassword        string
	JWTSecret            string
	JWTRefreshSecret     string
	JWTExpiresIn         time.Duration
	JWTRefreshExpiresIn  time.Duration
	Environment          string
	Port                 string
	CircuitBreakerConfig CircuitBreakerConfig
}

type CircuitBreakerConfig struct {
	Timeout        time.Duration
	ErrorThreshold int
	ResetTimeout   time.Duration
}

func Load() *Config {
	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "10m"))
	jwtRefreshExpiresIn, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRES_IN", "168h"))

	cbTimeout, _ := time.ParseDuration(getEnv("CIRCUIT_BREAKER_TIMEOUT", "5s"))
	cbErrorThreshold, _ := strconv.Atoi(getEnv("CIRCUIT_BREAKER_ERROR_THRESHOLD", "5"))
	cbResetTimeout, _ := time.ParseDuration(getEnv("CIRCUIT_BREAKER_RESET_TIMEOUT", "30s"))

	return &Config{
		DatabaseURL:         getEnv("DATABASE_URL", "postgres://user:password@localhost/trading_platform?sslmode=disable"),
		RedisURL:            getEnv("REDIS_URL", "redis://localhost:6379"),
		RedisPassword:       getEnv("REDIS_PASSWORD", ""),
		JWTSecret:           getEnv("JWT_SECRET", "your-secret-key"),
		JWTRefreshSecret:    getEnv("JWT_REFRESH_SECRET", "your-refresh-secret-key"),
		JWTExpiresIn:        jwtExpiresIn,
		JWTRefreshExpiresIn: jwtRefreshExpiresIn,
		Environment:         getEnv("GO_ENV", "development"),
		Port:                getEnv("PORT", "8080"),
		CircuitBreakerConfig: CircuitBreakerConfig{
			Timeout:        cbTimeout,
			ErrorThreshold: cbErrorThreshold,
			ResetTimeout:   cbResetTimeout,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
