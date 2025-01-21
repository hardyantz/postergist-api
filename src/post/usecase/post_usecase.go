package post

import (
	domain "postergist-api/src/post/domain"
	repository "postergist-api/src/post/repository"
)

type postUseacase struct {
	postRepo repository.PostRepository
}

type PostUc interface {
	GetPosts() ([]domain.Post, error)
	CreatePost(domain.Post) error
}

func NewPostUsecase(p repository.PostRepository) PostUc {
	return &postUseacase{
		postRepo: p,
	}
}

func (p *postUseacase) GetPosts() ([]domain.Post, error) {
	return p.postRepo.GetPosts()
}

func (p *postUseacase) CreatePost(dp domain.Post) error {
	return p.postRepo.CreatePosts(dp)
}
