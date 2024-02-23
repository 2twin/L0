package app

import (
	"log"

	"github.com/2twin/L0/internal/config"
	"github.com/2twin/L0/internal/repository"
	orderRepository "github.com/2twin/L0/internal/repository/order"
	"github.com/2twin/L0/internal/service"
	orderService "github.com/2twin/L0/internal/service/order"
)

type serviceProvider struct {
	natsConfig      config.NatsConfig
	postgresConfig  config.PostgresConfig
	httpConfig      config.HttpConfig
	orderRepository repository.OrderRepository
	orderService    service.OrderService
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) NatsConfig() config.NatsConfig {
	if s.natsConfig == nil {
		cfg, err := config.NewNatsConfig()
		if err != nil {
			log.Fatalf("failed to get nats config: %s", err.Error())
		}

		s.natsConfig = cfg
	}
	return s.natsConfig
}

func (s *serviceProvider) PostgresConfig() config.PostgresConfig {
	if s.postgresConfig == nil {
		cfg, err := config.NewPostgresConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		s.postgresConfig = cfg
	}
	return s.postgresConfig
}

func (s *serviceProvider) HttpConfig() config.HttpConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHttpConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}
	return s.httpConfig
}

func (s *serviceProvider) OrderRepository() repository.OrderRepository {
	if s.orderRepository == nil {
		dbURL := s.PostgresConfig().Url()
		rep, err := orderRepository.NewRepository(dbURL)
		if err != nil {
			log.Fatalf("failed to get order repository: %s", err.Error())
		}
		s.orderRepository = rep
	}
	return s.orderRepository
}

func (s *serviceProvider) OrderService() service.OrderService {
	if s.orderService == nil {
		s.orderService = orderService.NewService(s.OrderRepository())
	}

	return s.orderService
}
