package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type OrderRepo interface {
	ListOrder(limit, skip int) ([]model.OrderResp, error)
	GetById(id int) (model.OrderResp, error)
	CreateOrder(order model.Order) (model.Order, error)
	SubTotalOrder(order model.Order) (model.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func (o *orderRepo) ListOrder(limit, skip int) ([]model.OrderResp, error) {
	var result []model.OrderResp
	o.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(skip).Limit(limit)
	}).Find(&result)
	return result, nil
}

func (o *orderRepo) GetById(id int) (model.OrderResp, error) {
	result := model.OrderResp{}
	o.db.First(&result, id)
	o.db.First(&result.Cashier)
	o.db.First(&result.PaymentType)
	return result, nil
}

func (o *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	data := o.db.Create(&order)
	if data.Error != nil {
		return order, data.Error
	}
	return order, nil
}
func (o *orderRepo) SubTotalOrder(order model.Order) (model.Order, error) {
	o.db.Create(&order)
	return order, nil
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}
