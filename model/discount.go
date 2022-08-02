package model

import (
	"time"
)

type Discount struct {
	DiscountId      uint      `gorm:"primaryKey"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       time.Time `json:"expiredAt"`
	ExpiredAtFormat string
	StringFormat    string
}

type DiscountResp struct {
	DiscountId      uint      `json:"discountId"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       time.Time `json:"expiredAt"`
	ExpiredAtFormat string    `json:"expiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
}

func (DiscountResp) TableName() string {
	return "discounts"
}

func NewDiscount(Qty int, Type string, Result int, ExpiredAt time.Time) Discount {
	return Discount{
		Qty:       Qty,
		Type:      Type,
		Result:    Result,
		ExpiredAt: ExpiredAt,
	}
}
