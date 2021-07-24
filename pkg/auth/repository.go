package auth

import (
	"errors"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadUserByEmail(string) (*entities.User, error)
	ReadUserByUsername(string) (*entities.User, error)
	CreateUser(*entities.User) error
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

func (r *repo) ReadUserByEmail(email string) (*entities.User, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		log.Println("error while parsing pg url")
		return &entities.User{}, errors.New("error on db connection")
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		log.Println(err)
	}

	return user, err
}

func (r *repo) ReadUserByUsername(username string) (*entities.User, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		log.Println("error while parsing pg url")
		return &entities.User{}, errors.New("error on db connection")
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("username = ?", username).Select()

	if err != nil {
		log.Println(err)
	}

	return user, err
}

func (r *repo) CreateUser(user *entities.User) error {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		log.Println("error while parsing pg url")
		return errors.New("error on db connection")
	}

	db := pg.Connect(opt)
	defer db.Close()

	_, err = db.Model(user).
		Insert()

	if err != nil {
		log.Println(err)
	}

	return err
}
