package repository

import (
	"go-pos/model"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	ListPayment(limit, skip, subtotal int) []model.PaymentResp
	GetById(id int) (model.PaymentResp, error)
	CreatePayment(payment model.Payment) (model.NewPaymentResp, error)
	UpdatePayment(payment model.Payment, id int) error
	DeletePayment(id int) error
}

type paymentRepo struct {
	db *gorm.DB
}

func (p *paymentRepo) ListPayment(limit, skip, subtotal int) []model.PaymentResp {
	var result []model.PaymentResp
	p.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("subtotal = ?", subtotal).Offset(skip).Limit(limit)
	})
	return result
}

func (p *paymentRepo) GetById(id int) (model.PaymentResp, error) {
	result := model.PaymentResp{}
	err := p.db.First(&result, id)
	if err != nil {
		return result, err.Error
	}
	return result, nil
}

func (p *paymentRepo) CreatePayment(payment model.Payment) (model.NewPaymentResp, error) {
	var data model.NewPaymentResp
	err := p.db.Create(&payment)
	p.db.Find(&data, payment)
	if err != nil {
		return data, err.Error
	}
	return data, nil
}

func (p *paymentRepo) UpdatePayment(payment model.Payment, id int) error {
	err := p.db.Model(&payment).Where("id = ?", id).Updates(model.Payment{
		Name: payment.Name,
		Type: payment.Type,
		Logo: payment.Logo,
	})
	if (model.Payment{} == payment) {
		return err.Error
	}
	if err != nil {
		return err.Error
	}
	return nil
}

func (p *paymentRepo) DeletePayment(id int) error {
	var payment model.Payment
	err := p.db.First(&payment, id)
	if (model.Payment{} == payment) {
		return err.Error
	}
	err = p.db.Delete(&payment, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db: db}
}
