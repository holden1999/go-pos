package model

import "gorm.io/gorm"

type Payment struct {
	name string
	tipe string `db:"type"`
	logo string
	gorm.Model
}

func (p *Payment) getName() string {
	return p.name
}

func (p *Payment) getType() string {
	return p.tipe
}

func (p *Payment) getLogo() string {
	return p.logo
}

func (p *Payment) setName(code string) {
	p.name = code
}

func (p *Payment) setType(code string) {
	p.tipe = code
}

func (p *Payment) setLogo(code string) {
	p.logo = code
}

func NewPayment(name string, tipe string, logo string) Payment {
	return Payment{
		name: name,
		tipe: tipe,
		logo: logo,
	}
}
