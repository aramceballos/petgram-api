package auth

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	_ "github.com/lib/pq"
)

type Repository interface {
	ReadUserByEmail(string) (*entities.User, error)
	ReadUserByUsername(string) (*entities.User, error)
	CreateUser(*entities.User) error
}

type repo struct {
	db sql.DB
}

var postgresRepo *repo

func NewPostgresRepository() Repository {
	url := os.Getenv("DATABASE_URL")

	if postgresRepo == nil {
		db, err := sql.Open("postgres", url)
		if err != nil {
			log.Fatal("error connecting to db")
		}

		postgresRepo = &repo{
			db: *db,
		}
	}

	return postgresRepo
}

func (r *repo) ReadUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Name, &user.Username, &user.Password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return &user, fmt.Errorf("user not found")
	}

	return &user, err
}

func (r *repo) ReadUserByUsername(username string) (*entities.User, error) {
	var user entities.User

	err := r.db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.ID, &user.Email, &user.Name, &user.Username, &user.Password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return &user, fmt.Errorf("user not found")
	}

	return &user, err
}

func (r *repo) CreateUser(user *entities.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return err
}
