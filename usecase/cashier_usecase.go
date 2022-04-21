package usecase

import (
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type CashierUseCase interface {
	ListCashier(limit, skip int) []model.Cashier
	DetailCashier(id int) model.Cashier
	CreateCashier(cashier apprequest.Cashier) (model.Cashier, error)
	UpdateCashier(cashier apprequest.Cashier, id int) error
	DeleteCashier(id int) error
}

type cashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (c cashierUseCase) ListCashier(limit, skip int) []model.Cashier {
	return c.cashierRepo.ListCashier(limit, skip)
}

func (c cashierUseCase) DetailCashier(id int) model.Cashier {
	return c.cashierRepo.GetById(id)
}

func (c cashierUseCase) CreateCashier(cashier apprequest.Cashier) (model.Cashier, error) {
	newCashier := model.NewCashier(cashier.Name, cashier.Passcode)
	return c.cashierRepo.CreateCashier(newCashier)
}

func (c cashierUseCase) UpdateCashier(cashier apprequest.Cashier, id int) error {
	newCashier := model.NewCashier(cashier.Name, cashier.Passcode)
	return c.cashierRepo.UpdateCashier(newCashier, id)
}

func (c cashierUseCase) DeleteCashier(id int) error {
	return c.cashierRepo.DeleteCashier(id)
}

func NewCashierUseCase(cashierRepo repository.CashierRepo) CashierUseCase {
	return &cashierUseCase{cashierRepo: cashierRepo}
}
