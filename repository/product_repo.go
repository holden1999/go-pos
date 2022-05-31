package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	ListProduct(limit, skip, categoryId int) []model.Product
	GetById(id int) model.Product
	CreateProduct(product model.Product) (model.Product, error)
	UpdateProduct(product model.Product, id int) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *gorm.DB
}

func (p productRepo) ListProduct(limit, skip, categoryId int) []model.Product {
	var result []model.Product
	p.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("category = ?", categoryId).Offset(skip).Limit(limit)
	}).Find(&result)
	return result
}

func (p productRepo) GetById(id int) model.Product {
	result := model.Product{}
	p.db.First(&result, id)
	return result
}

func (p *productRepo) CreateProduct(product model.Product) (model.Product, error) {
	data := p.db.Create(&product)
	if data.Error != nil {
		return product, data.Error
	}
	return product, nil
}

func (p productRepo) UpdateProduct(product model.Product, id int) error {
	p.db.First(&product, id)
	p.db.Save(&product)
	return nil
}

func (p productRepo) DeleteProduct(id int) error {
	var product model.Product
	err := p.db.Delete(&product, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}
