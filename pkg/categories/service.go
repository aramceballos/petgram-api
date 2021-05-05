package categories

var categoriesRepo = NewPostgresRepository()

type CategoriesService interface {
	FindAll() (categories []Category, err error)
	Find(id int) (category Category, err error)
}

type service struct{}

func NewCategoriesService() CategoriesService {
	return &service{}
}

func (*service) FindAll() (categories []Category, err error) {
	categories, err = categoriesRepo.FindAll()

	return categories, err
}

func (*service) Find(id int) (category Category, err error) {
	category, err = categoriesRepo.Find(id)

	return category, err
}
