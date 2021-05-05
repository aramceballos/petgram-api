package categories

var categoriesRepo = NewPostgresRepository()

type CategoriesService interface {
	Find(id int) (category Category, err error)
}

type service struct{}

func NewCategoriesService() CategoriesService {
	return &service{}
}

func (*service) Find(id int) (category Category, err error) {
	category, err = categoriesRepo.Find(id)

	return category, err
}
