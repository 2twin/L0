package natsstreaming

import (
	"context"
	"log"

	"github.com/2twin/L0/internal/model"
	"github.com/2twin/L0/internal/repository"
	"github.com/nats-io/stan.go"
)

type natsStreaming struct {
	clusterID string
	clientID  string
	subject   string
	sc        stan.Conn
	addr      string
}

type NatsStreaming interface {
	Connect() error
	Publish(order *model.Order) error
	Subscribe(ctx context.Context, repo repository.OrderRepository) (stan.Subscription, error)
}

func NewNatsStreaming(clusterID string, clientID string, subject string, addr string) *natsStreaming {
	return &natsStreaming{
		clusterID: clusterID,
		clientID:  clientID,
		subject:   subject,
		addr:      addr,
	}
}

func (ns *natsStreaming) Connect() error {
	if ns.sc == nil {
		sc, err := stan.Connect(ns.clusterID, ns.clientID, stan.NatsURL(ns.addr))
		if err != nil {
			log.Printf("=====err: %v", err)
			return err
		}
		ns.sc = sc
	}

	return nil
}
