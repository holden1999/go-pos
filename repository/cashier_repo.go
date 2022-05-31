package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type CashierRepo interface {
	ListCashier(limit, skip int) []model.Cashier
	GetById(id int) model.Cashier
	CreateCashier(cashier model.Cashier) (model.Cashier, error)
	UpdateCashier(cashier model.Cashier, id int) error
	DeleteCashier(id int) error
}

type cashierRepo struct {
	db *gorm.DB
}

func (c cashierRepo) ListCashier(limit, skip int) []model.Cashier {
	var result []model.Cashier
	c.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(skip).Limit(limit)
	}).Find(&result)
	return result
}

func (c cashierRepo) GetById(id int) model.Cashier {
	result := model.Cashier{}
	c.db.First(&result, id)
	return result
}

func (c cashierRepo) CreateCashier(cashier model.Cashier) (model.Cashier, error) {
	data := c.db.Create(&cashier)
	if data.Error != nil {
		return cashier, data.Error
	}
	return cashier, nil
}

func (c cashierRepo) UpdateCashier(cashier model.Cashier, id int) error {
	c.db.First(&cashier, id)
	c.db.Save(&cashier)
	return nil
}

func (c cashierRepo) DeleteCashier(id int) error {
	var cashier model.Cashier
	err := c.db.Delete(&cashier, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewCashierRepo(db *gorm.DB) CashierRepo {
	return &cashierRepo{db: db}
}
