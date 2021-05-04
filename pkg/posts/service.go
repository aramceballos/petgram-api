package posts

var postsRepo = NewPostgresRepository()

type PostsService interface {
	Find(id int) (post Post, err error)
}

type service struct{}

func NewPostsService() PostsService {
	return &service{}
}

func (*service) Find(id int) (post Post, err error) {
	post, err = postsRepo.Find(id)

	return post, err
}
