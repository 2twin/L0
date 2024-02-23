package order

import "github.com/2twin/L0/internal/repository"

type service struct {
	orderRepository repository.OrderRepository
}

func NewService(orderRepository repository.OrderRepository) *service {
	return &service{
		orderRepository: orderRepository,
	}
}