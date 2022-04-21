package usecase

import (
	"go-pos/model"
	"go-pos/repository"
)

type OrderUseCase interface {
}

type orderUseCase struct {
	orderRepo repository.OrderRepo
}

func (o orderUseCase) ListOrder(limit, skip int) model.Order {
	return o.orderRepo.ListOrder(limit, skip)
}

func (o orderUseCase) DetailOrder(id int) model.Order {
	return o.orderRepo.GetById(id)
}

func (o orderUseCase) CreateOrder() {

}

func NewOrderUseCase(orderRepo repository.OrderRepo) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}
