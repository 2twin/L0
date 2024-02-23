package natsstreaming

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

func (ns *natsStreaming) Subscribe() (stan.Subscription, error) {
	sub, err := ns.sc.Subscribe(ns.subject, func(m *stan.Msg) {
		fmt.Println(m)
	})

	if err != nil {
		return nil, err
	}

	return sub, nil
}