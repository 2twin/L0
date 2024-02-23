package order

import (
	"context"

	"github.com/2twin/L0/internal/model"
)

func (s *service) Get(ctx context.Context, orderUUID string) (*model.Order, error) {
	order, err := s.orderRepository.Get(ctx, orderUUID)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, model.ErrorOrderNotFound
	}

	return order, nil
}