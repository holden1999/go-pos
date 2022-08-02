package repository

import (
	"errors"
	"go-pos/model"
	"gorm.io/gorm"
)

type CashierRepo interface {
	ListCashier(limit, skip int) []model.CashierResp
	GetById(id int) (model.CashierResp, error)
	CreateCashier(cashier model.Cashier) (model.CreateCashierResp, error)
	UpdateCashier(cashier model.Cashier, id int) error
	DeleteCashier(id int) error
}

type cashierRepo struct {
	db *gorm.DB
}

func (c *cashierRepo) ListCashier(limit, skip int) []model.CashierResp {
	var result []model.CashierResp
	c.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(skip).Limit(limit)
	}).Find(&result)
	return result
}

func (c *cashierRepo) GetById(id int) (model.CashierResp, error) {
	result := model.CashierResp{}
	err := c.db.First(&result, id)
	if err != nil {
		return result, err.Error
	}
	return result, nil
}

func (c *cashierRepo) CreateCashier(cashier model.Cashier) (model.CreateCashierResp, error) {
	if cashier.Name == "" || cashier.Passcode == "" {
		return model.CreateCashierResp{}, errors.New("incomplete data")
	}
	var data model.CreateCashierResp
	err := c.db.Create(&cashier)
	c.db.Find(&data, cashier)
	if err != nil {
		return data, err.Error
	}
	return data, nil
}

func (c *cashierRepo) UpdateCashier(cashier model.Cashier, id int) error {
	err := c.db.Where("name = ?", cashier.Name).First(&cashier, id)
	if err != nil {
		return err.Error
	}
	c.db.Save(&cashier)
	return nil
}

func (c *cashierRepo) DeleteCashier(id int) error {
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
