package order

import (
	"context"

	"github.com/2twin/L0/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, order *model.Order) (string, error) {
	var validate *validator.Validate
	err := validate.Struct(order)

	if err != nil {
		return "", err
	}

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
