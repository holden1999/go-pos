package model

import (
	"gorm.io/gorm"
	"time"
)

type OrderData struct {
	List interface{} `json:"orders"`
	Meta `json:"meta"`
}

type OrderResp struct {
	OrderId       uint        `gorm:"id"`
	CashierId     CashierResp `gorm:"embedded" gorm:"column:id" json:"cashierId"`
	PaymentTypeId PaymentResp `gorm:"embedded" gorm:"column:id" json:"paymentTypeId"`
	TotalPrice    string      `json:"totalPrice"`
	TotalPaid     string      `json:"totalPaid"`
	TotalReturn   string      `json:"totalReturn"`
	ReceiptId     string      `json:"receiptId"`
	CreatedAt     time.Time   `json:"createdAt"`
	Cashier       CashierResp `gorm:"embedded" json:"cashier"`
}

type ListOrder struct {
	OrderResp
	OrderPaymentResp
}

type Order struct {
	gorm.Model
	totalPrice  int64
	totalPaid   int64
	totalReturn int64
	product     []Product
	cashier     Cashier
	paymentType Payment
}

func (OrderResp) TableName() string {
	return "orders"
}

func NewOrder(totalPrice, totalPaid, totalReturn int64, product []Product, cashier Cashier, payment Payment) Order {
	return Order{
		totalPrice:  totalPrice,
		totalPaid:   totalPaid,
		totalReturn: totalReturn,
		product:     product,
		cashier:     cashier,
		paymentType: payment,
	}
}
