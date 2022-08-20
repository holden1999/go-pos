package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	ListProduct(limit, skip, categoryId int) []model.ProductResp
	GetById(id int) model.ProductResp
	CreateProduct(product model.Product, discount model.Discount) (model.NewProductResp, error)
	UpdateProduct(product model.Product, id int) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *gorm.DB
}

func (p *productRepo) ListProduct(limit, skip, categoryId int) []model.ProductResp {
	var result []model.ProductResp
	p.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryId).Offset(skip).Limit(limit)
	}).Find(&result)
	return result
}

func (p *productRepo) GetById(id int) model.ProductResp {
	result := model.ProductResp{}
	p.db.First(&result, id)
	p.db.First(&result.Category)
	p.db.First(&result.Discount)
	return result
}

func (p *productRepo) CreateProduct(product model.Product, discount model.Discount) (model.NewProductResp, error) {
	var data model.NewProductResp
	err := p.db.Create(&product)
	err = p.db.Create(&discount)
	p.db.Find(&data, product)
	if err != nil {
		return data, err.Error
	}
	return data, nil
}

func (p *productRepo) UpdateProduct(product model.Product, id int) error {
	err := p.db.First(&product, id)
	if err != nil {
		return err.Error
	}
	p.db.Save(&product)
	return nil
}

func (p *productRepo) DeleteProduct(id int) error {
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
