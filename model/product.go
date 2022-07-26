package model

import (
	"gorm.io/gorm"
	"time"
)

type ProductData struct {
	Products interface{} `json:"products"`
	Meta     List        `json:"meta"`
}

type Product struct {
	*gorm.Model
	Sku      string
	Name     string
	Stock    int
	Price    int
	Image    string
	category Category
	discount Discount
}

type NewProductResp struct {
	ProductId  uint      `gorm:"primarykey" gorm:"column:id" json:"productId"`
	CategoryId int       `gorm:"column:category" json:"categoryId"`
	Name       string    `json:"name"`
	Sku        string    `json:"sku"`
	Image      string    `json:"image"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ProductResp struct {
	ProductId uint         `gorm:"column:id" json:"productId"`
	Sku       string       `json:"sku"`
	Name      string       `json:"name"`
	Stock     int          `json:"stock"`
	Price     int          `json:"price"`
	Image     string       `json:"image"`
	Category  CategoryResp `gorm:"embedded" json:"category"`
	Discount  DiscountResp `gorm:"embedded" json:"discount"`
}

func (ProductResp) TableName() string {
	return "products"
}

func (NewProductResp) TableName() string {
	return "products"
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
