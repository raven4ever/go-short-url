package db

import (
	"github.com/redis/go-redis/v9"
	"short-url/config"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.EnvConfig.RedisConfig.GetAddress(),
		Username: config.EnvConfig.RedisConfig.Username,
		Password: config.EnvConfig.RedisConfig.Password,
		DB:       config.EnvConfig.RedisConfig.DB,
	})
}
