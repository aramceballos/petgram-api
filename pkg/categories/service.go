package categories

import "github.com/aramceballos/petgram-api/pkg/entities"

type Service interface {
	FetchCategories() ([]entities.Category, error)
	FetchCategory(int) (entities.Category, error)
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

func (s *service) FetchCategories() ([]entities.Category, error) {
	categories, err := s.repository.ReadCategories()

	return categories, err
}

func (s *service) FetchCategory(id int) (entities.Category, error) {
	category, err := s.repository.ReadCategory(id)

	return category, err
}
