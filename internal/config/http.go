package config

import (
	"errors"
	"os"
)

const (
	httpPort = "HTTP_PORT"
)

type httpConfig struct {
	port string
}

type HttpConfig interface {
	Address() string
}

func NewHttpConfig() (*httpConfig, error) {
	port := os.Getenv(httpPort)
	if port == "" {
		return nil, errors.New("HTTP_PORT environment variable is not set")
	}

	return &httpConfig{
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return ":" + cfg.port
}