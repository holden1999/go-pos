package main

import (
	"go-pos/config"
	"go-pos/delivery/api"
	"go-pos/manager"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.InfraManager
	usecase manager.UseCaseManager
}

func (s *server) Run() {
	cashierApiGroup := s.config.RouterEngine.Group("/cashiers")
	categoryApiGroup := s.config.RouterEngine.Group("/categories")
	paymentApiGroup := s.config.RouterEngine.Group("/payments")
	productApiGroup := s.config.RouterEngine.Group("/products")
	orderApiGroup := s.config.RouterEngine.Group("/orders")
	reportApiGroup := s.config.RouterEngine.Group("/")
	api.NewCashierApi(cashierApiGroup)
	api.NewCategoryApi(categoryApiGroup)
	api.NewPaymentApi(paymentApiGroup)
	api.NewProductApi(productApiGroup, s.usecase.ProductUseCase())
	api.NewOrderApi(orderApiGroup)
	api.NewOrderApi(reportApiGroup)
	err := s.config.RouterEngine.Run(s.config.ApiUrl)
	if err != nil {
		return
	}
}

func NewApiServer() Server {
	newConfig := config.NewConfig()
	infra := manager.NewInfra(newConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &server{
		config:  newConfig,
		infra:   infra,
		usecase: usecase,
	}
}
