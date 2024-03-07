package main

import (
	"context"
	"fmt"
	"log"
	"short-url/db"

	"github.com/redis/go-redis/v9"
)

func main() {
	log.Println("Starting the activity...")

	log.Println("Connecting to Redis...")
	ctx := context.Background()

	err := db.RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := db.RedisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := db.RedisClient.Get(ctx, "key2").Result()
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
