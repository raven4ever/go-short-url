package main

import (
	"context"
	"log"
	"short-url/api/server"
	"short-url/config"
	"short-url/db"
)

func main() {
	log.Println("Starting the activity...")

	log.Println("Connecting to Redis...")
	redisClient := db.NewRedisClient()

	// test the Redis connection
	ctx := context.Background()
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	srv := server.NewApiServer(config.EnvConfig.ServerConfig.Addr(), redisClient)

	// Start server
	log.Println("Starting the server...")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
