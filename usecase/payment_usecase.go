package usecase

import (
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type PaymentUseCase interface {
	ListPayment(limit, skip, subtotal int) []model.Payment
	DetailPayment(id int) model.Payment
	CreatePayment(payment apprequest.Payment) (model.Payment, error)
	UpdatePayment(payment apprequest.Payment, id int) error
	DeletePayment(id int) error
}

type paymentUseCase struct {
	paymentRepo repository.PaymentRepo
}

func (p paymentUseCase) ListPayment(limit, skip, subtotal int) []model.Payment {
	return p.paymentRepo.ListPayment(limit, skip, subtotal)
}

func (p paymentUseCase) DetailPayment(id int) model.Payment {
	return p.paymentRepo.GetById(id)
}

func (p paymentUseCase) CreatePayment(payment apprequest.Payment) (model.Payment, error) {
	newPayment := model.NewPayment(payment.Name, payment.Type, payment.Logo)
	return p.paymentRepo.CreatePayment(newPayment)
}

func (p paymentUseCase) UpdatePayment(payment apprequest.Payment, id int) error {
	newPayment := model.NewPayment(payment.Name, payment.Type, payment.Logo)
	return p.paymentRepo.UpdatePayment(newPayment, id)
}

func (p paymentUseCase) DeletePayment(id int) error {
	return p.paymentRepo.DeletePayment(id)
}

func NewPaymentUseCase(paymentRepo repository.PaymentRepo) PaymentUseCase {
	return &paymentUseCase{paymentRepo: paymentRepo}
}
