package usecase

import (
	"go-pos/controller/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type OrderUseCase interface {
	ListOrder(limit, skip int) ([]model.OrderResp, error)
	DetailOrder(id int) (model.OrderResp, error)
	CreateOrder(order apprequest.Order) (model.Order, error)
	SubTotalOrder(order apprequest.Order) (model.Order, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepo
}

func (o *orderUseCase) ListOrder(limit, skip int) ([]model.OrderResp, error) {
	return o.orderRepo.ListOrder(limit, skip)
}

func (o *orderUseCase) DetailOrder(id int) (model.OrderResp, error) {
	return o.orderRepo.GetById(id)
}

func (o *orderUseCase) CreateOrder(order apprequest.Order) (model.Order, error) {
	newOrder := model.NewOrder(order.TotalPaid)
	return o.orderRepo.CreateOrder(newOrder)
}

func (o *orderUseCase) SubTotalOrder(order apprequest.Order) (model.Order, error) {
	newOrder := model.NewOrder(order.TotalPaid)
	return o.orderRepo.SubTotalOrder(newOrder)
}

func NewOrderUseCase(orderRepo repository.OrderRepo) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}
