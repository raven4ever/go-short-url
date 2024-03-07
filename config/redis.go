package config

import "fmt"

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
	TTL      int
}

func (rc *RedisConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", rc.Host, rc.Port)
}
