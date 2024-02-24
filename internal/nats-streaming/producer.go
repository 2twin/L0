package natsstreaming

import (
	"encoding/json"

	"github.com/2twin/L0/internal/model"
)

func (ns *natsStreaming) Publish(order *model.Order) error {
	orderJson, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ns.sc.Publish(ns.subject, []byte(orderJson))

	if err != nil {
		return err
	}

	return nil
}
