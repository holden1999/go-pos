package model

import (
	"gorm.io/gorm"
	"time"
)

type CashierData struct {
	Cashiers interface{} `json:"cashiers"`
	Meta     `json:"meta"`
}

type Cashier struct {
	gorm.Model
	Name     string
	Passcode string
}

type CashierResp struct {
	CashierId uint   `gorm:"column:id" json:"cashierId"`
	Name      string `json:"name"`
}

type CreateCashierResp struct {
	Passcode  string    `json:"passcode"`
	CashierId uint      `gorm:"column:id" json:"cashierId"`
	Name      string    `json:"name" json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (CashierResp) TableName() string {
	return "Cashiers"
}

func (CreateCashierResp) TableName() string {
	return "Cashiers"
}

func NewCashier(name string, passcode string) Cashier {
	return Cashier{
		Name:     name,
		Passcode: passcode,
	}
}
