package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string
}

type CategoryData struct {
	Category interface{} `json:"categories"`
	Meta     List        `json:"meta"`
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
