package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductData struct {
	Products []ProductResp `json:"products"`
	Meta     `json:"meta"`
}

type Product struct {
	ProductId  uint           `gorm:"primaryKey" json:"productId"`
	Sku        string         `json:"sku"`
	Name       string         `json:"name"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	Image      string         `json:"image"`
	CategoryId uint           `json:"categoryId"`
	DiscountId uint           `json:"discountId"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type NewProductResp struct {
	ProductId  uint      `json:"productId"`
	CategoryId uint      `json:"categoryId"`
	Name       string    `json:"name"`
	Sku        string    `json:"sku"`
	Image      string    `json:"image"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ProductOrder struct {
	ProductId uint `json:"productId"`
	Qty       uint `json:"qty"`
}

type ProductResp struct {
	ProductId uint                `json:"productId"`
	Sku       string              `json:"sku"`
	Name      string              `json:"name"`
	Stock     int                 `json:"stock"`
	Price     int                 `json:"price"`
	Image     string              `json:"image"`
	Category  CategoryProductResp `gorm:"embedded" json:"category"`
	Discount  DiscountResp        `gorm:"embedded" json:"discount"`
}

func (ProductResp) TableName() string {
	return "products"
}

func (NewProductResp) TableName() string {
	return "products"
}

func NewProduct(categoryId uint, name string, image string, stock int, price int) Product {
	return Product{
		CategoryId: categoryId,
		Name:       name,
		Image:      image,
		Stock:      stock,
		Price:      price,
	}
}
