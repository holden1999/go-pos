package model

import "gorm.io/gorm"

type Cashier struct {
	gorm.Model
	Name     string
	Passcode string
}

func NewCashier(name string, passcode string) Cashier {
	return Cashier{
		Name:     name,
		Passcode: passcode,
	}
}
