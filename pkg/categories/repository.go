package categories

import (
	"errors"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadCategories() ([]entities.Category, error)
	ReadCategory(int) (entities.Category, error)
}

type repo struct {
	url string
}

var postgresRepo *repo

func NewPostgresRepository() Repository {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	if postgresRepo == nil {
		postgresRepo = &repo{
			url: url,
		}
	}

	return postgresRepo
}

func (r *repo) ReadCategories() ([]entities.Category, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		log.Println("error while parsing pg url")
		return []entities.Category{}, errors.New("error on db connection")
	}

	db := pg.Connect(opt)
	defer db.Close()

	var c []entities.Category
	err = db.Model(&c).Select()

	if err != nil {
		log.Println(err)
	}

	return c, err
}

func (r *repo) ReadCategory(id int) (entities.Category, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		log.Println("error while parsing pg url")
		return entities.Category{}, errors.New("error on db connection")
	}

	db := pg.Connect(opt)
	defer db.Close()

	c := &entities.Category{ID: id}
	err = db.Model(c).WherePK().Select()

	if err != nil {
		log.Println(err)
	}

	return *c, err
}
