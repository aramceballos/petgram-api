package users

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
)

type Service interface {
	FetchUser(string) (entities.User, error)
	FetchUserById(int) (entities.User, error)
}

type service struct {
	repository Repository
}

func NewService() Service {
	postgresRepository := NewPostgresRepository()

	return &service{
		repository: postgresRepository,
	}
}

func (s *service) FetchUser(username string) (entities.User, error) {
	user, err := s.repository.ReadUser(username)

	return user, err
}

func (s *service) FetchUserById(id int) (entities.User, error) {
	user, err := s.repository.ReadUserById(id)

	return user, err
}
