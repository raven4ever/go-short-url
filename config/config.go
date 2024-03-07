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
			Port:     stringToInt(getEnvVariable("REDIS_PORT", "6379")),
			Username: getEnvVariable("REDIS_USERNAME", ""),
			Password: getEnvVariable("REDIS_PASSWORD", ""),
			DB:       stringToInt(getEnvVariable("REDIS_DB", "0")),
			TTL:      time.Duration(stringToInt(getEnvVariable("REDIS_TTL", "5"))) * time.Minute,
		},
		ServerConfig: ServerConfig{
			Port: stringToInt(getEnvVariable("PORT", "8080")),
		},
		ShortUrlLength: stringToInt(getEnvVariable("SHORT_URL_LENGTH", "7")),
	}
}

func getEnvVariable(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func stringToInt(intValue string) int {
	intPort, err := strconv.Atoi(intValue)
	if err != nil {
		log.Fatalf("%s is not a valid int number", intValue)
	}
	return intPort
}
