package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type OrderRepo interface {
	ListOrder(limit, skip int) model.Order
	GetById(id int) model.Order
	CreateOrder(order model.Order) (model.Order, error)
	SubTotalOrder(order model.Order) model.Order
}

type orderRepo struct {
	db *gorm.DB
}

func (o orderRepo) ListOrder(limit, skip int) model.Order {
	result := model.Order{}
	o.db.Find(&result)
	return result
}

func (o orderRepo) GetById(id int) model.Order {
	result := model.Order{}
	o.db.First(&result, id)
	return result
}

func (o orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	data := o.db.Create(&order)
	if data.Error != nil {
		return order, data.Error
	}
	return order, nil
}
func (o orderRepo) SubTotalOrder(order model.Order) model.Order {
	o.db.Create(&order)
	return order
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}
