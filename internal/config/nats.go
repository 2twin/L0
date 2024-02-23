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
	clusterID string
	clientID  string
	subject   string
}

type NatsConfig interface {
	ClusterID() string
	ClientID() string
	Subject() string
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
		clusterID: cluster,
		clientID:  client,
		subject:   subject,
	}, nil
}

func (cfg *natsConfig) ClientID() string {
	return cfg.clientID
}

func (cfg *natsConfig) ClusterID() string {
	return cfg.clusterID
}

func (cfg *natsConfig) Subject() string {
	return cfg.subject
}
