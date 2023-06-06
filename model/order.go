package model

import "time"

type ListOrder struct {
	List []OrderResp `json:"orders"`
	Meta `json:"meta"`
}

type DetailOrder struct {
	Order OrderResp `json:"order"`
}

type AddOrder struct {
	Order    `json:"order"`
	Products []ProductOrderResp `json:"products"`
}

type Order struct {
	OrderId        uint      `gorm:"primaryKey" json:"orderId"`
	CashiersId     uint      `json:"cashiersId"`
	PaymentTypesId uint      `json:"paymentTypesId"`
	TotalPrice     int64     `json:"totalPrice"`
	TotalPaid      int64     `json:"totalPaid"`
	TotalReturn    int64     `json:"totalReturn"`
	ReceiptId      uint      `json:"receiptId"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}

type OrderDetail struct {
	OrderId   uint
	ProductId uint
	Quantity  int
}

type OrderResp struct {
	OrderId        uint             `json:"orderId"`
	CashiersId     uint             `json:"cashiersId"`
	PaymentTypesId uint             `json:"paymentTypesId"`
	TotalPrice     int64            `json:"totalPrice"`
	TotalPaid      int64            `json:"totalPaid"`
	TotalReturn    int64            `json:"totalReturn"`
	ReceiptId      uint             `json:"receiptId"`
	CreatedAt      time.Time        `json:"createdAt"`
	Cashier        CashierResp      `gorm:"embedded" json:"Cashier"`
	PaymentType    OrderPaymentResp `gorm:"embedded" json:"payment_type"`
}

func (OrderDetail) TableName() string {
	return "orders"
}

func (OrderResp) TableName() string {
	return "orders"
}

func NewOrder(totalPaid int64) Order {
	return Order{
		TotalPaid: totalPaid,
	}
}
