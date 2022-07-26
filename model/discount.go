package model

import (
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	*gorm.Model
	Qty             int
	Type            string
	Result          int
	ExpiredAt       time.Time
	ExpiredAtFormat string
	StringFormat    string
}

type DiscountResp struct {
	DiscountId      uint      `gorm:"column:id" json:"discountId"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       time.Time `json:"expiredAt"`
	ExpiredAtFormat string    `json:"expiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
}
