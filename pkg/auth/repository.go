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

func (r *repo) ReadUserByEmail(email string) (*entities.User, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return &entities.User{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("email = ?", email).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return user, err
}

func (r *repo) ReadUserByUsername(username string) (*entities.User, error) {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return &entities.User{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	user := &entities.User{}
	err = db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return user, err
}

func (r *repo) CreateUser(user *entities.User) error {
	opt, err := pg.ParseURL(r.url)
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
