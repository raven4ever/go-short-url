package api

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type ApiServer struct {
	Address     string
	RedisClient *redis.Client
}

func NewApiServer(address string, client *redis.Client) *ApiServer {
	return &ApiServer{
		Address:     address,
		RedisClient: client,
	}
}

func (s *ApiServer) Run() error {
	echoServer := echo.New()

	// create the API v1 group
	apiV1Group := echoServer.Group("/api/v1")

	// create the URL shortening endpoint
	apiV1Group.POST("/shorten", nil)

	// start the server
	return echoServer.Start(s.Address)
}
