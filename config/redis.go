package config

import (
	"fmt"
	"time"
)

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
	TTL      time.Duration
}

func (rc *RedisConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", rc.Host, rc.Port)
}
