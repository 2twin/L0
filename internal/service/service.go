package service

import (
	"context"

	"github.com/2twin/L0/internal/model"
)

type OrderService interface {
	Create(ctx context.Context, order *model.Order) (string, error)
	Get(ctx context.Context, orderUUID string) (*model.Order, error)
}
