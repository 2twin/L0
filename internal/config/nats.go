package config

import (
	"errors"
	"os"
)

const (
	natsClusterId = "NATS_CLUSTER_ID"
	natsClientId  = "NATS_CLIENT_ID"
	natsSubject   = "NATS_SUBJECT"
)

type natsConfig struct {
	ClusterID string
	ClientID  string
	Subject   string
}

type NatsConfig interface {
	Address() string
}

func NewNatsConfig() (*natsConfig, error) {
	cluster := os.Getenv(natsClusterId)
	if cluster == "" {
		return nil, errors.New("NATS_CLUSTER_ID environment variable is not set")
	}

	client := os.Getenv(natsClientId)
	if client == "" {
		return nil, errors.New("NATS_CLIENT_ID environment variable is not set")
	}

	subject := os.Getenv(natsSubject)
	if subject == "" {
		return nil, errors.New("NATS_SUBJECT environment variable is not set")
	}

	return &natsConfig{
		ClusterID: cluster,
		ClientID:  client,
		Subject:   subject,
	}, nil
}

func (cfg *natsConfig) Address() string {
	return "nats://nats-streming:4222"
}