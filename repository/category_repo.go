package repository

import (
	"go-pos/model"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	ListCategory(limit, skip int) []model.CategoryResp
	GetById(id int) model.CategoryResp
	CreateCategory(category model.Category) (model.Category, error)
	UpdateCategory(category model.Category, id int) error
	DeleteCategory(id int) error
}

type categoryRepo struct {
	db *gorm.DB
}

func (c categoryRepo) ListCategory(limit, skip int) []model.CategoryResp {
	var result []model.CategoryResp
	c.db.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(skip).Limit(limit)
	}).Find(&result)
	return result
}

func (c categoryRepo) GetById(id int) model.CategoryResp {
	result := model.CategoryResp{}
	c.db.First(&result, id)
	return result
}

func (c categoryRepo) CreateCategory(category model.Category) (model.Category, error) {
	c.db.Create(&category)
	return category, nil
}

func (c categoryRepo) UpdateCategory(category model.Category, id int) error {
	c.db.First(&category, id)
	c.db.Save(&category)
	return nil
}

func (c categoryRepo) DeleteCategory(id int) error {
	var category model.Category
	err := c.db.Delete(&category, id)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepo{db: db}
}
