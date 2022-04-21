package manager

import "go-pos/usecase"

type UseCaseManager interface {
	CashierUseCase() usecase.CashierUseCase
	CategoryUseCase() usecase.CategoryUseCase
	ProductUseCase() usecase.ProductUseCase
	PaymentUseCase() usecase.PaymentUseCase
	OrderUseCase() usecase.OrderUseCase
	ReportUseCase() usecase.ReportUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CashierUseCase() usecase.CashierUseCase {
	return usecase.NewCashierUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) CategoryUseCase() usecase.CategoryUseCase {
	return usecase.NewCategoryUseCase(u.repo.CategoryRepo())
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repo.ProductRepo())
}

func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repo.PaymentRepo())
}

func (u *useCaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(u.repo.OrderRepo())
}

func (u *useCaseManager) ReportUseCase() usecase.ReportUseCase {
	return usecase.NewReportUseCase(u.repo.ReportRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
