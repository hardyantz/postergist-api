package category

import (
	domain "postergist-api/src/domain/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	GetCategories() ([]domain.Category, error)
	CreateCategories(domain.Category) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		DB: db,
	}
}

func (c *categoryRepository) GetCategories() ([]domain.Category, error) {
	var categories []domain.Category

	result := c.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (p *categoryRepository) CreateCategories(dp domain.Category) error {
	tx := p.DB.Create(dp)
	return tx.Error
}
