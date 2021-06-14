package users

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadUser(string) (entities.User, error)
}

type repo struct {
	url string
}

func NewPostgresRepository() Repository {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	return &repo{
		url: url,
	}
}

func (r *repo) ReadUser(username string) (entities.User, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return entities.User{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	u := &entities.User{Username: username}
	err = db.Model(u).
		Where("username = ?", u.Username).
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return *u, err
}
