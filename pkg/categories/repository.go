package categories

type Reader interface {
	FindAll() (categories []Category, err error)
	Find(id int) (category Category, err error)
}

type Repository interface {
	Reader
}
