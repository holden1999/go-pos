package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	CategoryId uint           `gorm:"primaryKey" json:"categoryId"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type CategoryData struct {
	Category interface{} `json:"categories"`
	Meta     `json:"meta"`
}

type CategoryResp struct {
	CategoryId int       `json:"categoryId"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updatedAt"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CategoryProductResp struct {
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
}

func (CategoryProductResp) TableName() string {
	return "categories"
}

func (CategoryResp) TableName() string {
	return "categories"
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}
