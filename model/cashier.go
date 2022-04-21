package model

import "gorm.io/gorm"

type Cashier struct {
	gorm.Model
	name     string
	passcode string
}

func (c *Cashier) getName() string {
	return c.name
}

func (c *Cashier) getPasscode() string {
	return c.passcode
}

func (c *Cashier) setName(code string) {
	c.name = code
}

func (c *Cashier) setPasscode(code string) {
	c.passcode = code
}

func NewCashier(name string, passcode string) Cashier {
	return Cashier{
		name:     name,
		passcode: passcode,
	}
}
