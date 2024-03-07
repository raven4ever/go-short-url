package db

import "github.com/redis/go-redis/v9"
import "short-url/config"

var RedisClient = newRedisClient()

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.EnvConfig.RedisConfig.GetAddress(),
		Username: config.EnvConfig.RedisConfig.Username,
		Password: config.EnvConfig.RedisConfig.Password,
		DB:       config.EnvConfig.RedisConfig.DB,
	})
}
