package usecase

import (
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type OrderUseCase interface {
	ListOrder(limit, skip int) []model.OrderResp
	DetailOrder(id int) model.Order
	CreateOrder(order apprequest.Order) (model.Order, error)
	SubTotalOrder(order apprequest.Order) model.Order
}

type orderUseCase struct {
	orderRepo repository.OrderRepo
}

func (o orderUseCase) ListOrder(limit, skip int) []model.OrderResp {
	return o.orderRepo.ListOrder(limit, skip)
}

func (o orderUseCase) DetailOrder(id int) model.Order {
	return o.orderRepo.GetById(id)
}

func (o orderUseCase) CreateOrder(order apprequest.Order) (model.Order, error) {
	newOrder := model.NewOrder(order.TotalPrice, order.TotalPaid, order.TotalReturn, order.Product, order.Cashier, order.PaymentType)
	return o.orderRepo.CreateOrder(newOrder)
}

func (o orderUseCase) SubTotalOrder(order apprequest.Order) model.Order {
	newOrder := model.NewOrder(order.TotalPrice, order.TotalPaid, order.TotalReturn, order.Product, order.Cashier, order.PaymentType)
	return o.orderRepo.SubTotalOrder(newOrder)
}

func NewOrderUseCase(orderRepo repository.OrderRepo) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}
