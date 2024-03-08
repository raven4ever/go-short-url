package endpoints

import (
	"context"
	"net/http"
	"short-url/config"
	"short-url/utils"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type EndpointsService struct {
	RedisClient *redis.Client
}

type ModalData struct {
	ShortURL string
	Err      error
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

	// validate the URL
	if !utils.IsValidURL(url) {
		return c.String(http.StatusBadRequest, "Invalid URL!")
	}

	// generate a random string using alphanumeric characters
	short := utils.RandString(config.EnvConfig.ShortUrlLength)

	// store the URL in Redis
	ctx := context.Background()
	err := s.RedisClient.Set(ctx, short, url, config.EnvConfig.RedisConfig.TTL).Err()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "add_error.html", ModalData{Err: err})
	}

	return c.Render(http.StatusOK, "modal.html", ModalData{ShortURL: short})
}

func (s *EndpointsService) RedirectURL(c echo.Context) error {
	shortURL := c.Param("shortURL")

	// retrieve the URL from Redis
	ctx := context.Background()
	url, err := s.RedisClient.Get(ctx, shortURL).Result()
	if err != nil {
		return c.String(http.StatusNotFound, "Key not found or already expired!")
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)
}
