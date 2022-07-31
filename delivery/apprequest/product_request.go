package apprequest

import "go-pos/model"

type ProductRequest struct {
	CategoryId model.Category `json:"category_id"`
	Name       string         `json:"name"`
	Image      string         `json:"image"`
	Stock      int            `json:"stock"`
	Price      int            `json:"price"`
	Discount   Discount       `json:"discount"`
}
