package users

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
)

type Service interface {
	FetchUser(string) (entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchUser(username string) (entities.User, error) {
	user, err := s.repository.ReadUser(username)

	return user, err
}
