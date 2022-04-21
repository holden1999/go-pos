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
