package categories

type Reader interface {
	Find(id int) (category Category, err error)
}

type Repository interface {
	Reader
}
