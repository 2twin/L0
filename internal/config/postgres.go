package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	postgresUser     = "POSTGRES_USER"
	postgresPassword = "POSTGRES_PASSWORD"
	postgresHost     = "POSTGRES_HOST"
	postgresPort     = "POSTGRES_PORT"
	postgresDB       = "POSTGRES_DB"
)

type postgresConfig struct {
	user     string
	password string
	host     string
	port     string
	db       string
}

type PostgresConfig interface {
	Url() string
}

func NewPostgresConfig() (*postgresConfig, error) {
	user := os.Getenv(postgresUser)
	if user == "" {
		return nil, errors.New("POSTGRES_USER environment variable is not set")
	}

	password := os.Getenv(postgresPassword)
	if user == "" {
		return nil, errors.New("POSTGRES_PASSWORD environment variable is not set")
	}

	host := os.Getenv(postgresHost)
	if user == "" {
		return nil, errors.New("POSTGRES_HOST environment variable is not set")
	}

	port := os.Getenv(postgresPort)
	if user == "" {
		return nil, errors.New("POSTGRES_PORT environment variable is not set")
	}

	db := os.Getenv(postgresDB)
	if user == "" {
		return nil, errors.New("POSTGRES_DB environment variable is not set")
	}

	return &postgresConfig{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		db:       db,
	}, nil
}

func (cfg *postgresConfig) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.user, cfg.password, cfg.host, cfg.port, cfg.db)
}
