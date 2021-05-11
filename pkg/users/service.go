package users

import "github.com/aramceballos/petgram-api/pkg/entities"

type Service interface {
	InsertUser(*entities.User) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertUser(user *entities.User) error {
	err := s.repository.CreateUser(user)

	return err
}
