package usecase

import (
	"go-pos/delivery/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type ProductUseCase interface {
	ListProduct(limit int, skip int, categoryId int) []model.Product
	DetailProduct(id int) model.Product
	CreateProduct(product apprequest.ProductRequest) (model.Product, error)
	UpdateProduct(product apprequest.ProductRequest, id int) error
	DeleteProduct(id int) error
}

type productUseCase struct {
	productRepo repository.ProductRepo
}

func (p productUseCase) ListProduct(limit, skip, categoryId int) []model.Product {
	return p.productRepo.ListProduct(limit, skip, categoryId)
}

func (p *productUseCase) DetailProduct(id int) model.Product {
	return p.productRepo.GetById(id)
}

func (p *productUseCase) CreateProduct(product apprequest.ProductRequest) (model.Product, error) {
	newProduct := model.NewProduct(product.Sku, product.Name, product.Image, product.Stock, product.Price, product.Discount, product.CategoryId)
	return p.productRepo.CreateProduct(newProduct)
}

func (p productUseCase) UpdateProduct(product apprequest.ProductRequest, id int) error {
	newProduct := model.NewProduct(product.Sku, product.Name, product.Image, product.Stock, product.Price, product.Discount, product.CategoryId)
	return p.productRepo.UpdateProduct(newProduct, id)
}

func (p productUseCase) DeleteProduct(id int) error {
	return p.productRepo.DeleteProduct(id)
}

func NewProductUseCase(productRepo repository.ProductRepo) ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}
