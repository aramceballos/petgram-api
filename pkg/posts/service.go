package posts

var postsRepo = NewPostgresRepository()

type PostsService interface {
	FindAll() (posts []Post, err error)
	Find(id int) (post Post, err error)
}

type service struct{}

func NewPostsService() PostsService {
	return &service{}
}

func (*service) FindAll() (posts []Post, err error) {
	posts, err = postsRepo.FindAll()

	return posts, err
}

func (*service) Find(id int) (post Post, err error) {
	post, err = postsRepo.Find(id)

	return post, err
}
