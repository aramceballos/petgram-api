package posts

type Reader interface {
	Find(id int) (post Post, err error)
}

type Repository interface {
	Reader
}
