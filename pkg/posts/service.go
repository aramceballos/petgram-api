package posts

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
)

type Service interface {
	FetchPosts() ([]entities.Post, error)
	FetchPost(int) (entities.Post, error)
	LikePost(int, int) error
	UnlikePost(int, int) error
	FetchLikedPosts(int) ([]entities.Post, error)
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

func (s *service) LikePost(user_id int, post_id int) error {
	err := s.repository.LikePost(user_id, post_id)

	return err
}

func (s *service) UnlikePost(user_id int, post_id int) error {
	err := s.repository.UnlikePost(user_id, post_id)

	return err
}

func (s *service) FetchLikedPosts(user_id int) ([]entities.Post, error) {
	liked_posts, err := s.repository.ReadLikedPosts(user_id)

	return liked_posts, err
}
