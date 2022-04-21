package model

import (
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Sku      string
	Name     string
	Image    string
	Stock    int
	Price    int
	discount Discount
	category Category
}

func NewProduct(sku string, name string, image string, stock int, price int, discount Discount, category Category) Product {
	return Product{
		Sku:      sku,
		Name:     name,
		Image:    image,
		Stock:    stock,
		Price:    price,
		discount: discount,
		category: category,
	}
}
