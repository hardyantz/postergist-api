package post

import (
	domain "postergist-api/src/post/domain"
	repository "postergist-api/src/post/repository"
)

type postUsecase struct {
	postRepo repository.PostRepository
}

type PostUc interface {
	GetPosts() ([]domain.Post, error)
	CreatePost(domain.Post) error
	GetPost(id int) (domain.Post, error)
}

func NewPostUsecase(p repository.PostRepository) PostUc {
	return &postUsecase{
		postRepo: p,
	}
}

func (p *postUsecase) GetPosts() ([]domain.Post, error) {
	return p.postRepo.GetPosts()
}

func (p *postUsecase) CreatePost(dp domain.Post) error {
	return p.postRepo.CreatePosts(dp)
}

func (p *postUsecase) GetPost(id int) (domain.Post, error) {
	return p.postRepo.GetPost(id)
}
