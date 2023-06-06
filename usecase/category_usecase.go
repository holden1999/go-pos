package usecase

import (
	"go-pos/controller/apprequest"
	"go-pos/model"
	"go-pos/repository"
)

type CategoryUseCase interface {
	ListCategory(limit, skip int) []model.CategoryResp
	DetailCategory(id int) model.CategoryResp
	CreateCategory(category apprequest.CategoryRequest) (model.CategoryResp, error)
	UpdateCategory(category apprequest.CategoryRequest, id int) error
	DeleteCategory(id int) error
}

type categoryUseCase struct {
	categoryRepo repository.CategoryRepo
}

func (c categoryUseCase) ListCategory(limit, skip int) []model.CategoryResp {
	return c.categoryRepo.ListCategory(limit, skip)
}

func (c categoryUseCase) DetailCategory(id int) model.CategoryResp {
	return c.categoryRepo.GetById(id)
}

func (c categoryUseCase) CreateCategory(category apprequest.CategoryRequest) (model.CategoryResp, error) {
	newCategory := model.NewCategory(category.Name)
	return c.categoryRepo.CreateCategory(newCategory)
}

func (c categoryUseCase) UpdateCategory(category apprequest.CategoryRequest, id int) error {
	newCategory := model.NewCategory(category.Name)
	return c.categoryRepo.UpdateCategory(newCategory, id)
}

func (c categoryUseCase) DeleteCategory(id int) error {
	return c.categoryRepo.DeleteCategory(id)
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepo) CategoryUseCase {
	return &categoryUseCase{categoryRepo: categoryRepo}
}
