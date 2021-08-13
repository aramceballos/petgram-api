package categories

import (
	"log"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadCategories() ([]entities.Category, error)
	ReadCategory(int) (entities.Category, error)
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

func (r *repo) ReadCategories() ([]entities.Category, error) {
	var c []entities.Category
	err := r.db.Model(&c).Select()

	return c, err
}

func (r *repo) ReadCategory(id int) (entities.Category, error) {
	c := &entities.Category{ID: id}
	err := r.db.Model(c).WherePK().Select()

	return *c, err
}
