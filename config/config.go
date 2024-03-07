package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	RedisConfig    RedisConfig
	ServerConfig   ServerConfig
	ShortUrlLength int
}

var EnvConfig = loadEnvirontmentConfig()

func loadEnvirontmentConfig() *Config {
	return &Config{
		RedisConfig: RedisConfig{
			Host:     getEnvVariable("REDIS_HOST", "127.0.0.1"),
			Port:     stringToPort(getEnvVariable("REDIS_PORT", "6379")),
			Username: getEnvVariable("REDIS_USERNAME", ""),
			Password: getEnvVariable("REDIS_PASSWORD", ""),
			DB:       stringToPort(getEnvVariable("REDIS_DB", "0")),
			TTL:      time.Duration(stringToPort(getEnvVariable("REDIS_TTL", "5"))) * time.Minute,
		},
		ServerConfig: ServerConfig{
			Port: stringToPort(getEnvVariable("PORT", "8080")),
		},
		ShortUrlLength: stringToPort(getEnvVariable("SHORT_URL_LENGTH", "7")),
	}
}

func getEnvVariable(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func stringToPort(port string) int {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("%s is not a valid port number", port)
	}
	return intPort
}
