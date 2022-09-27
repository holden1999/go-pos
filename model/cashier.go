package model

import (
	"gorm.io/gorm"
	"time"
)

type CashierData struct {
	Cashiers []CashierResp `json:"cashiers"`
	Meta     `json:"meta"`
}

type Cashier struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Passcode  string
}

type CashierResp struct {
	CashierId uint   `gorm:"column:id" json:"cashierId"`
	Name      string `json:"name"`
}

type CreateCashierResp struct {
	Passcode  string    `json:"passcode"`
	CashierId uint      `gorm:"column:id" json:"cashierId"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (CashierResp) TableName() string {
	return "cashiers"
}

func (CreateCashierResp) TableName() string {
	return "cashiers"
}

func NewCashier(name string, passcode string) Cashier {
	return Cashier{
		Name:     name,
		Passcode: passcode,
	}
}
