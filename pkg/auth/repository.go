package auth

import (
	"log"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadUserByEmail(string) (*entities.User, error)
	ReadUserByUsername(string) (*entities.User, error)
	CreateUser(*entities.User) error
}

type repo struct {
	db pg.DB
}

var postgresRepo *repo

func NewPostgresRepository() Repository {
	url := "postgres://postgres:postgress@db:5432/postgres?sslmode=disable"

	if postgresRepo == nil {
		opt, err := pg.ParseURL(url)
		if err != nil {
			log.Fatal("error parsing db url")
		}
		db := pg.Connect(opt)

		postgresRepo = &repo{
			db: *db,
		}
	}

	return postgresRepo
}

func (r *repo) ReadUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.Model(user).Where("email = ?", email).Select()

	return user, err
}

func (r *repo) ReadUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.Model(user).Where("username = ?", username).Select()

	return user, err
}

func (r *repo) CreateUser(user *entities.User) error {
	_, err := r.db.Model(user).
		Insert()

	return err
}
