package server

import (
	"html/template"
	"short-url/api/v1"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// enable server logging
	echoServer.Use(middleware.Logger())

	// create the API v1 group
	apiV1Group := echoServer.Group("/api/v1")

	// create the URL shortening endpoint
	sh := endpoints.NewEndpointsService(s.RedisClient)
	sh.RegisterRoutes(apiV1Group)

	// add static middleware
	echoServer.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "static",
		Browse:     true,
		Index:      "index.html",
		HTML5:      true,
		IgnoreBase: true,
	}))

	// set the custom template renderer
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	echoServer.Renderer = t

	// start the server
	return echoServer.Start(s.Address)
}
