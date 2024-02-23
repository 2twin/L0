package natsstreaming

import "github.com/nats-io/stan.go"

type natsStreaming struct {
	clusterID string
	clientID  string
	subject   string
	sc        stan.Conn
}

type NatsStreaming interface {
	Connect() error
	Publish() error
	Subscribe() (stan.Subscription, error)
}

func NewNatsStreaming(clusterID string, clientID string, subject string) *natsStreaming {
	return &natsStreaming{
		clusterID: clusterID,
		clientID:  clientID,
		subject:   subject,
	}
}

func (ns *natsStreaming) Connect() error {
	if ns.sc == nil {
		sc, err := stan.Connect(ns.clusterID, ns.clientID)
		if err != nil {
			return err
		}
		ns.sc = sc
	}

	return nil
}
