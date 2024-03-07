package main

import (
	"context"
	"fmt"
	"log"
	"short-url/config"

	"github.com/redis/go-redis/v9"
)

func main() {
	log.Println("Starting the activity...")

	log.Println("Connecting to Redis...")
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.EnvConfig.RedisConfig.GetAddress(),
		Username: config.EnvConfig.RedisConfig.Username,
		Password: config.EnvConfig.RedisConfig.Password,
		DB:       config.EnvConfig.RedisConfig.DB,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
