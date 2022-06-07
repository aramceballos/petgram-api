package categories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	_ "github.com/lib/pq"
)

type Repository interface {
	ReadCategories() ([]entities.Category, error)
	ReadCategory(int) (entities.Category, error)
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

func (r *repo) ReadCategories() ([]entities.Category, error) {
	var categories []entities.Category

	stmt, err := r.db.Prepare("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.ID, &category.Category, &category.ImageURL)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, err
}

func (r *repo) ReadCategory(id int) (entities.Category, error) {
	var category entities.Category

	err := r.db.QueryRow("SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Category, &category.ImageURL)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return category, fmt.Errorf("category not found")
	}

	return category, err
}
