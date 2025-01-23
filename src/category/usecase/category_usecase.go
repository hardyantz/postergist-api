package category

import (
	domain "postergist-api/src/category/domain"
	repository "postergist-api/src/category/repository"
)

type categoryUsecase struct {
	catRepo repository.CategoryRepository
}

type CategoryUc interface {
	GetCategories() ([]domain.Category, error)
	CreateCategory(dc domain.Category) error
	GetCategory(id int) (domain.Category, error)
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

func (c *categoryUsecase) GetCategory(id int) (domain.Category, error) {
	return c.catRepo.GetCategory(id)
}
