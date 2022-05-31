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
	routerEngine   *gin.Engine
	config         *config.Config
	infra          manager.InfraManager
	useCaseManager manager.UseCaseManager
}

func (s *server) initHandlers() {
	s.v1()
}
func (s *server) v1() {
	//s.routerEngine.Use(middleware.AuthTokenMiddleware())
	categoryApiGroup := s.config.RouterEngine.Group("/categories")
	paymentApiGroup := s.config.RouterEngine.Group("/payments")
	productApiGroup := s.config.RouterEngine.Group("/products")
	orderApiGroup := s.config.RouterEngine.Group("/orders")
	reportApiGroup := s.config.RouterEngine.Group("/")
	api.NewCategoryApi(categoryApiGroup, s.useCaseManager.CategoryUseCase())
	api.NewPaymentApi(paymentApiGroup, s.useCaseManager.PaymentUseCase())
	api.NewProductApi(productApiGroup, s.useCaseManager.ProductUseCase())
	api.NewOrderApi(orderApiGroup, s.useCaseManager.OrderUseCase())
	api.NewReportApi(reportApiGroup, s.useCaseManager.ReportUseCase())

}

func (s *server) v2() {
	cashierApiGroup := s.config.RouterEngine.Group("/cashiers")
	api.NewCashierApi(cashierApiGroup, s.useCaseManager.CashierUseCase())
}

func (s *server) Run() {
	s.initHandlers()
	err := s.config.RouterEngine.Run(s.config.ApiUrl)
	if err != nil {
		return
	}
}

func NewApiServer() Server {
	newConfig := config.NewConfig()
	infra := manager.NewInfra(newConfig)
	repo := manager.NewRepoManager(infra)
	UseCaseManager := manager.NewUseCaseManager(repo)
	return &server{
		config:         newConfig,
		infra:          infra,
		useCaseManager: UseCaseManager,
	}
}
