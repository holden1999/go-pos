package main

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/delivery/api"
	"go-pos/manager"
)

type Server interface {
	Run()
}

type server struct {
	routerEngine *gin.Engine
	config       *config.Config
	infra        manager.InfraManager
	usecase      manager.UseCaseManager
}

func (s *server) initHandlers() {

}
func (s *server) v1() {
	//s.routerEngine.Use(middleware.AuthTokenMiddleware())
	categoryApiGroup := s.config.RouterEngine.Group("/categories")
	paymentApiGroup := s.config.RouterEngine.Group("/payments")
	productApiGroup := s.config.RouterEngine.Group("/products")
	orderApiGroup := s.config.RouterEngine.Group("/orders")
	reportApiGroup := s.config.RouterEngine.Group("/")
	api.NewCategoryApi(categoryApiGroup, s.usecase.CategoryUseCase())
	api.NewPaymentApi(paymentApiGroup, s.usecase.PaymentUseCase())
	api.NewProductApi(productApiGroup, s.usecase.ProductUseCase())
	api.NewOrderApi(orderApiGroup, s.usecase.OrderUseCase())
	api.NewReportApi(reportApiGroup, s.usecase.ReportUseCase())

}

func (s *server) v2() {
	cashierApiGroup := s.config.RouterEngine.Group("/cashiers")
	api.NewCashierApi(cashierApiGroup, s.usecase.CashierUseCase())
}

func (s *server) Run() {

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
