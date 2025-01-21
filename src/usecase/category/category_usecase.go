package category

import (
	domain "postergist-api/src/domain/category"
	repository "postergist-api/src/repository/category"
)

type categoryUsecase struct {
	catRepo repository.CategoryRepository
}

type CategoryUc interface {
	GetCategories() ([]domain.Category, error)
	CreateCategory(dc domain.Category) error
}

func NewPostUsecase(c repository.CategoryRepository) CategoryUc {
	return &categoryUsecase{
		catRepo: c,
	}
}

func (c *categoryUsecase) GetCategories() ([]domain.Category, error) {
	return c.catRepo.GetCategories()
}

func (c *categoryUsecase) CreateCategory(dc domain.Category) error {
	return c.catRepo.CreateCategories(dc)
}
