package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type PaymentRepo interface {
	ListPayment(limit, skip, subtotal int) []model.PaymentResp
	GetById(id int) model.Payment
	CreatePayment(payment model.Payment) (model.Payment, error)
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

func (p *paymentRepo) GetById(id int) model.Payment {
	result := model.Payment{}
	p.db.First(&result, id)
	return result
}

func (p *paymentRepo) CreatePayment(payment model.Payment) (model.Payment, error) {
	data := p.db.Create(&payment)
	if data.Error != nil {
		return payment, data.Error
	}
	return payment, nil
}

func (p *paymentRepo) UpdatePayment(payment model.Payment, id int) error {
	err := p.db.First(&payment, id)
	if err != nil {
		return err.Error
	}
	p.db.Save(&payment)
	return nil
}

func (p *paymentRepo) DeletePayment(id int) error {
	var payment model.Payment
	err := p.db.Delete(&payment, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db: db}
}
