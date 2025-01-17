package post

import (
	domain "postergist-api/src/domain/post"
	repository "postergist-api/src/repository/post"
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
