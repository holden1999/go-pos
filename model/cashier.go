package model

import "gorm.io/gorm"

type Cashier struct {
	gorm.Model
	Name     string
	Passcode string
}

func (c *Cashier) getName() string {
	return c.Name
}

func (c *Cashier) getPasscode() string {
	return c.Passcode
}

func (c *Cashier) setName(code string) {
	c.Name = code
}

func (c *Cashier) setPasscode(code string) {
	c.Passcode = code
}

func NewCashier(name string, passcode string) Cashier {
	return Cashier{
		Name:     name,
		Passcode: passcode,
	}
}
