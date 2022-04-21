package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	totalPrice  int64
	totalPaid   int64
	totalReturn int64
	product     []Product
	cashier     Cashier
	paymentType Payment
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
