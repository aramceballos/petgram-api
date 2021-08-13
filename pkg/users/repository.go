package users

import (
	"log"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadUser(string) (entities.User, error)
	ReadUserById(int) (entities.User, error)
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

func (r *repo) ReadUser(username string) (entities.User, error) {
	u := &entities.User{Username: username}
	err := r.db.Model(u).
		Where("username = ?", u.Username).
		Select()

	return *u, err
}

func (r *repo) ReadUserById(id int) (entities.User, error) {
	u := &entities.User{ID: id}
	err := r.db.Model(u).
		WherePK().
		Select()

	return *u, err
}
