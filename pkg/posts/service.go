package posts

import "github.com/aramceballos/petgram-api/pkg/entities"

type Service interface {
	FetchPosts() ([]entities.Post, error)
	FetchPost(int) (entities.Post, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchPosts() ([]entities.Post, error) {
	posts, err := s.repository.ReadPosts()

	return posts, err
}

func (s *service) FetchPost(id int) (entities.Post, error) {
	post, err := s.repository.ReadPost(id)

	return post, err
}
