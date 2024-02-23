package order

import (
	"context"

	"github.com/2twin/L0/internal/model"
	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, order *model.Order) (string, error) {
	orderUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	err = s.orderRepository.Create(ctx, orderUUID.String(), order)
	if err != nil {
		return "", err
	}

	return orderUUID.String(), nil
}