package users

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
)

type Service interface {
	FetchUser(int) (entities.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) FetchUser(id int) (entities.User, error) {
	user, err := s.repository.ReadUser(id)

	return user, err
}
