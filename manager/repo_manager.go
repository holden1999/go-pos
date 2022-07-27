package manager

import "go-pos/repository"

type RepoManager interface {
	AuthRepo() repository.AuthenticationRepo
	CashierRepo() repository.CashierRepo
	CategoryRepo() repository.CategoryRepo
	PaymentRepo() repository.PaymentRepo
	ProductRepo() repository.ProductRepo
	OrderRepo() repository.OrderRepo
	ReportRepo() repository.ReportRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) AuthRepo() repository.AuthenticationRepo {
	return repository.NewAuthRepo(r.infra.SqlDb())
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.NewCashierRepo(r.infra.SqlDb())
}

func (r *repoManager) CategoryRepo() repository.CategoryRepo {
	return repository.NewCategoryRepo(r.infra.SqlDb())
}

func (r *repoManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepo(r.infra.SqlDb())
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(r.infra.SqlDb())
}

func (r *repoManager) OrderRepo() repository.OrderRepo {
	return repository.NewOrderRepo(r.infra.SqlDb())
}

func (r *repoManager) ReportRepo() repository.ReportRepo {
	return repository.NewReportRepo(r.infra.SqlDb())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repoManager{manager}
}
