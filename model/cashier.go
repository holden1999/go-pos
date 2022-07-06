package model

import "gorm.io/gorm"

type Cashier struct {
	gorm.Model
	Name     string
	Passcode string
}

type CashierPasscode struct {
	Passcode string
}

type CashierToken struct {
	Token string
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
