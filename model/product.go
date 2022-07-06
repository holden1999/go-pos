package model

import (
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	productId uint `gorm:"primarykey"`
	Sku       string
	Name      string
	Stock     int
	Price     int
	Image     string
	category  Category
	discount  Discount
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
