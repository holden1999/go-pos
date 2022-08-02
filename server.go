package main

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/controller/api"
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
	s.public()
	s.private()
}
func (s *server) public() {
	cashierApiGroup := s.config.RouterEngine.Group("/cashiers")
	orderApiGroup := s.config.RouterEngine.Group("/orders")
	reportApiGroup := s.config.RouterEngine.Group("/")
	api.NewCashierApi(cashierApiGroup, s.useCaseManager.CashierUseCase())
	api.NewOrderApi(orderApiGroup, s.useCaseManager.OrderUseCase())
	api.NewReportApi(reportApiGroup, s.useCaseManager.ReportUseCase())
}

func (s *server) private() {
	productApiGroup := s.config.RouterEngine.Group("/products")
	categoryApiGroup := s.config.RouterEngine.Group("/categories")
	paymentApiGroup := s.config.RouterEngine.Group("/payments")
	LoginApiGroup := s.config.RouterEngine.Group("/cashiers")
	api.NewProductApi(productApiGroup, s.useCaseManager.ProductUseCase())
	api.NewCategoryApi(categoryApiGroup, s.useCaseManager.CategoryUseCase())
	api.NewPaymentApi(paymentApiGroup, s.useCaseManager.PaymentUseCase())
	api.NewLoginApi(LoginApiGroup, s.useCaseManager.LoginUseCase(), s.useCaseManager.JwtUseCase())

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
