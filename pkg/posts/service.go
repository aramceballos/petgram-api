package posts

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
)

type Service interface {
	FetchPosts() ([]entities.Post, error)
	FetchPostsByUserID(int) ([]entities.Post, error)
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

func (s *service) FetchPostsByUserID(userId int) ([]entities.Post, error) {
	posts, err := s.repository.ReadPostsByUserID(userId)

	return posts, err
}

func (s *service) FetchPost(id int) (entities.Post, error) {
	post, err := s.repository.ReadPost(id)

	return post, err
}

func (s *service) LikePost(userId int, postId int) error {
	err := s.repository.LikePost(userId, postId)

	return err
}

func (s *service) UnlikePost(userId int, postId int) error {
	err := s.repository.UnlikePost(userId, postId)

	return err
}

func (s *service) FetchLikedPosts(userId int) ([]entities.Post, error) {
	likedPosts, err := s.repository.ReadLikedPosts(userId)

	return likedPosts, err
}
