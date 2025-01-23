package post

import (
	"fmt"
	"postergist-api/helper"
	categoryRepository "postergist-api/src/category/repository"
	domain "postergist-api/src/post/domain"

	"gorm.io/gorm"
)

type postRepository struct {
	DB      *gorm.DB
	catRepo categoryRepository.CategoryRepository
}

type PostRepository interface {
	GetPosts() ([]domain.Post, error)
	CreatePosts(domain.Post) error
	GetPost(id int) (domain.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	catRepository := categoryRepository.NewCategoryRepository(db)
	return &postRepository{
		DB:      db,
		catRepo: catRepository,
	}
}

func (p *postRepository) GetPosts() ([]domain.Post, error) {
	var posts []domain.Post

	result := p.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (p *postRepository) CreatePosts(dp domain.Post) error {
	dp.CreatedDate = helper.GetCurrentDateTime()
	isCatExists, err := p.catRepo.GetCategory(dp.CategoryID)
	if err != nil {
		return fmt.Errorf("category not found")
	}
	if isCatExists.ID == 0 {
		return fmt.Errorf("category id is null")
	}
	tx := p.DB.Create(&dp)
	return tx.Error
}

func (p *postRepository) GetPost(id int) (domain.Post, error) {
	var post domain.Post
	tx := p.DB.First(&post, id)
	return post, tx.Error
}
