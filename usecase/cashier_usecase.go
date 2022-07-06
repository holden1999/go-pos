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

	GetPasscode(id int) model.CashierPasscode
	VerifyLogin(id int, cashierLogin apprequest.CashierLogin) (model.CashierToken, error)
	VerifyLogout(id int, cashierLogout apprequest.CashierLogin) error
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

func (c cashierUseCase) GetPasscode(id int) model.CashierPasscode {
	//TODO implement me
	panic("implement me")
}

func (c cashierUseCase) VerifyLogin(id int, cashier apprequest.CashierLogin) (model.CashierToken, error) {
	data := model.NewCashierPasscode(cashier.Passcode)
	return c.cashierRepo.VerifyLogin(id, data)
}

func (c cashierUseCase) VerifyLogout(id int, cashierLogout apprequest.CashierLogin) error {
	//TODO implement me
	panic("implement me")
}

func NewCashierUseCase(cashierRepo repository.CashierRepo) CashierUseCase {
	return &cashierUseCase{cashierRepo: cashierRepo}
}
