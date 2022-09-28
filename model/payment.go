package model

import (
	"gorm.io/gorm"
	"time"
)

type PaymentData struct {
	Payment []PaymentResp `json:"payments"`
	Meta    `json:"meta"`
}

type NewPaymentResp struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Logo      string    `json:"logo"`
	UpdateAt  time.Time `json:"updateAt"`
	CreatedAt time.Time `json:"createdAt"`
	PaymentId uint      `json:"paymentId"`
}

type PaymentResp struct {
	PaymentId uint   `gorm:"column:id" json:"paymentId"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Logo      string `json:"logo"`
}

type OrderPaymentResp struct {
	PaymentTypeId uint   `gorm:"column:id" json:"paymentTypeId"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
	Type          string `json:"type"`
}

type Payment struct {
	PaymentId uint           `gorm:"primaryKey" json:"paymentId"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Logo      string         `json:"logo"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updateAt"`
	DeleteAt  gorm.DeletedAt `gorm:"index" json:"deleteAt"`
}

func (NewPaymentResp) TableName() string {
	return "payments"
}

func NewPayment(name string, tipe string, logo string) Payment {
	return Payment{
		Name: name,
		Type: tipe,
		Logo: logo,
	}
}
