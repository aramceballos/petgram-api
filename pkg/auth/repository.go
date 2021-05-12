package auth

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadUserByEmail(string) (*entities.User, error)
	ReadUserByUsername(string) (*entities.User, error)
	CreateUser(*entities.User) error
}

type repo struct{}

var instance *repo

func NewPostgresRepository() Repository {
	if instance == nil {
		instance = &repo{}
	}

	return instance
}

func (*repo) ReadUserByEmail(e string) (*entities.User, error) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return &entities.User{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("email = ?", e).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return user, err
}

func (*repo) ReadUserByUsername(u string) (*entities.User, error) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return &entities.User{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("username = ?", u).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return user, err
}

func (*repo) CreateUser(user *entities.User) error {

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return err
	}

	db := pg.Connect(opt)
	defer db.Close()
	result, err := db.Model(user).
		Insert()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	return err
}
