package categories

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

type repo struct{}

var instance *repo

func NewPostgresRepository() Repository {
	if instance == nil {
		instance = &repo{}
	}

	return instance
}

func (*repo) Find(id int) (category Category, err error) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return
	}

	db := pg.Connect(opt)
	defer db.Close()

	c := &Category{ID: id}
	err = db.Model(c).WherePK().Select()

	return *c, err
}
