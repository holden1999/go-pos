package model

import "gorm.io/gorm"

type PaymentData struct {
	Payment []PaymentResp `json:"payments"`
	Meta    `json:"meta"`
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
	Name string
	Type string
	Logo string
	gorm.Model
}

func NewPayment(name string, tipe string, logo string) Payment {
	return Payment{
		Name: name,
		Type: tipe,
		Logo: logo,
	}
}
