package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	CategoryId uint `gorm:"primarykey"`
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type CategoryData struct {
	Category interface{} `json:"categories"`
	Meta     `json:"meta"`
}

type CategoryResp struct {
	CategoryId uint   `gorm:"column:id" json:"categoryId"`
	Name       string `json:"name"`
}

func (CategoryResp) TableName() string {
	return "categories"
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}
