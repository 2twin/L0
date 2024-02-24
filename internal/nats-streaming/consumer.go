package natsstreaming

import (
	"context"
	"encoding/json"
	"log"

	"github.com/2twin/L0/internal/model"
	"github.com/2twin/L0/internal/repository"
	"github.com/nats-io/stan.go"
)

func (ns *natsStreaming) Subscribe(ctx context.Context, repo repository.OrderRepository) (stan.Subscription, error) {
	sub, err := ns.sc.Subscribe(ns.subject, func(m *stan.Msg) {
		var order model.Order

		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println(err)
			return
		}

		err = repo.Create(ctx, order.OrderUID, &order)
		if err != nil {
			log.Println(err)
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return sub, nil
}
