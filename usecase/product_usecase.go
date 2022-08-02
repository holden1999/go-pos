package usecase

import (
	"go-pos/controller/apprequest"
	"go-pos/model"
	"go-pos/repository"
	"time"
)

type ProductUseCase interface {
	ListProduct(limit int, skip int, categoryId int) []model.ProductResp
	DetailProduct(id int) model.ProductResp
	CreateProduct(product apprequest.ProductRequest) (model.NewProductResp, error)
	UpdateProduct(product apprequest.ProductRequest, id int) error
	DeleteProduct(id int) error
}

type productUseCase struct {
	productRepo repository.ProductRepo
}

func (p productUseCase) ListProduct(limit, skip, categoryId int) []model.ProductResp {
	return p.productRepo.ListProduct(limit, skip, categoryId)
}

func (p *productUseCase) DetailProduct(id int) model.ProductResp {
	return p.productRepo.GetById(id)
}

func (p *productUseCase) CreateProduct(product apprequest.ProductRequest) (model.NewProductResp, error) {
	timeResult := time.Unix(int64(product.Discount.ExpiredAt), 0)
	newProduct := model.NewProduct(product.CategoryId, product.Name, product.Image, product.Stock, product.Price)
	newDiscount := model.NewDiscount(product.Discount.Qty, product.Discount.Type, product.Discount.Result, timeResult)
	return p.productRepo.CreateProduct(newProduct, newDiscount)
}

func (p productUseCase) UpdateProduct(product apprequest.ProductRequest, id int) error {
	newProduct := model.NewProduct(product.CategoryId, product.Name, product.Image, product.Stock, product.Price)
	return p.productRepo.UpdateProduct(newProduct, id)
}

func (p productUseCase) DeleteProduct(id int) error {
	return p.productRepo.DeleteProduct(id)
}

func NewProductUseCase(productRepo repository.ProductRepo) ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}
