package apprequest

import "go-pos/model"

type Order struct {
	TotalPrice  int64           `json:"totalPrice"`
	TotalPaid   int64           `json:"totalPaid"`
	TotalReturn int64           `json:"totalReturn"`
	Product     []model.Product `json:"product"`
	Cashier     model.Cashier   `json:"cashier"`
	PaymentType model.Payment   `json:"paymentType"`
}
