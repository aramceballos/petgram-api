package posts

type Reader interface {
	FindAll() (posts []Post, err error)
	Find(id int) (post Post, err error)
}

type Repository interface {
	Reader
}
