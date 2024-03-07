package endpoints

import (
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

	return c.String(200, url)
}

func (s *EndpointsService) RedirectURL(c echo.Context) error {
	shortURL := c.Param("shortURL")

	return c.String(200, shortURL)
}
