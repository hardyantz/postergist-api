package post

import (
	"postergist-api/helper"
	domain "postergist-api/src/post/domain"

	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

type PostRepository interface {
	GetPosts() ([]domain.Post, error)
	CreatePosts(domain.Post) error
	GetPost(id int) (domain.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		DB: db,
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
	tx := p.DB.Create(&dp)
	return tx.Error
}

func (p *postRepository) GetPost(id int) (domain.Post, error) {
	var post domain.Post
	tx := p.DB.First(&post, id)
	return post, tx.Error
}
