package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	_ "github.com/lib/pq"
)

type Repository interface {
	ReadUser(string) (entities.User, error)
	ReadUserById(int) (entities.User, error)
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

func (r *repo) ReadUser(username string) (entities.User, error) {
	var user entities.User

	err := r.db.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.ID, &user.Email, &user.Name, &user.Username, &user.Password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return user, fmt.Errorf("user not found")
	}

	return user, err
}

func (r *repo) ReadUserById(id int) (entities.User, error) {
	var user entities.User

	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Name, &user.Username, &user.Password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return user, fmt.Errorf("user not found")
	}

	return user, err
}
