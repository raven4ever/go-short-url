package endpoints

import (
	"context"
	"net/http"
	"short-url/config"
	"short-url/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type EndpointsService struct {
	RedisClient *redis.Client
}

func NewEndpointsService(client *redis.Client) *EndpointsService {
	return &EndpointsService{
		RedisClient: client,
	}
}

func (s *EndpointsService) RegisterRoutes(g *echo.Group) {
	// create the URL shortening endpoint
	g.POST("/shorten", s.ShortenURL)
	// create the URL redirection endpoint
	g.GET("/:shortURL", s.RedirectURL)
}

func (s *EndpointsService) ShortenURL(c echo.Context) error {
	url := c.FormValue("url")

	// generate a random string using alphanumeric characters
	short := utils.RandString(config.EnvConfig.ShortUrlLength)

	// store the URL in Redis
	ctx := context.Background()
	err := s.RedisClient.Set(ctx, short, url, time.Duration(config.EnvConfig.RedisConfig.TTL)).Err()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, short)
}

func (s *EndpointsService) RedirectURL(c echo.Context) error {
	shortURL := c.Param("shortURL")

	return c.String(200, shortURL)
}
