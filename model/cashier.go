package model

import (
	"gorm.io/gorm"
	"time"
)

type CashierData struct {
	Cashiers interface{} `json:"cashiers"`
	Meta     List        `json:"meta"`
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
	Passcode  string    `json:"passcode,omitempty"`
	CashierId uint      `gorm:"column:id" json:"cashierId" json:"cashierId,omitempty"`
	Name      string    `json:"name" json:"name,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type CashierPasscode struct {
	Passcode string
}

type CashierToken struct {
	Token string
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

func NewCashierPasscode(passcode string) CashierPasscode {
	return CashierPasscode{Passcode: passcode}
}

func NewCashierToken(token string) CashierToken {
	return CashierToken{Token: token}
}
