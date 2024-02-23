package repository

import (
	"context"

	"github.com/2twin/L0/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, orderUUID string, order *model.Order) error
	Get(ctx context.Context, orderUUID string) (*model.Order, error)
}
