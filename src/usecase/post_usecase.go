package usecase

import "postergist-api/src/repository"

type postUseacase struct {
	postRepo repository.PostRepository
}

type PostUc interface {
	GetPosts() error
}

func NewPostUsecase(p repository.PostRepository) PostUc {
	return &postUseacase{
		postRepo: p,
	}
}

func (p *postUseacase) GetPosts() error {
	return nil
}
