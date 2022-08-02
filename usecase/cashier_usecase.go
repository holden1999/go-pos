package usecase

import (
	"go-pos/controller/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type CashierUseCase interface {
	ListCashier(limit, skip int) []model.CashierResp
	DetailCashier(id int) (model.CashierResp, error)
	CreateCashier(cashier apprequest.Cashier) (model.CreateCashierResp, error)
	UpdateCashier(cashier apprequest.Cashier, id int) error
	DeleteCashier(id int) error
}

type cashierUseCase struct {
	cashierRepo repository.CashierRepo
}

func (c cashierUseCase) ListCashier(limit, skip int) []model.CashierResp {
	return c.cashierRepo.ListCashier(limit, skip)
}

func (c cashierUseCase) DetailCashier(id int) (model.CashierResp, error) {
	return c.cashierRepo.GetById(id)
}

func (c cashierUseCase) CreateCashier(cashier apprequest.Cashier) (model.CreateCashierResp, error) {
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
