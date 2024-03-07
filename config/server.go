package config

import (
	"fmt"
)

type ServerConfig struct {
	Port int
}

func (sc *ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", sc.Port)
}
