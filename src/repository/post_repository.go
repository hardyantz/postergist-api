package repository

type postRepository struct {
}

type PostRepository interface {
	GetPosts() error
}

func NewPostRepository() PostRepository {
	return &postRepository{}
}

func (p *postRepository) GetPosts() error {
	return nil
}
