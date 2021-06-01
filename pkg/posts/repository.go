package posts

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadPosts() ([]entities.Post, error)
	ReadPost(int) (entities.Post, error)
	LikePost(int, int) error
	UnlikePost(int, int) error
	ReadLikedPosts(int) ([]entities.Post, error)
}

type repo struct{}

var instance *repo

func NewPostgresRepository() Repository {
	if instance == nil {
		instance = &repo{}
	}

	return instance
}

func (*repo) ReadPosts() ([]entities.Post, error) {

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return []entities.Post{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	var p []entities.Post
	err = db.Model(&p).
		ColumnExpr("post.*").
		ColumnExpr("u.name, u.username, u.email").
		Join("JOIN users AS u ON u.id = post.user_id").
		Relation("Likes").
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return p, err
}

func (*repo) ReadPost(id int) (entities.Post, error) {

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return entities.Post{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	p := &entities.Post{ID: id}
	err = db.Model(p).
		ColumnExpr("post.*").
		ColumnExpr("u.name, u.username, u.email").
		WherePK().
		Join("JOIN users AS u ON u.id = post.user_id").
		Relation("Likes").
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return *p, err
}

func (*repo) LikePost(user_id int, post_id int) error {

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

	l := &entities.Like{UserID: user_id, PostID: post_id}
	_, err = db.Model(l).Insert()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (*repo) UnlikePost(user_id int, post_id int) error {

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

	l := &entities.Like{}
	_, err = db.Model(l).Where("user_id = ? AND post_id = ?", user_id, post_id).Delete()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (*repo) ReadLikedPosts(user_id int) ([]entities.Post, error) {

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	opt, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return []entities.Post{}, err
	}

	db := pg.Connect(opt)
	defer db.Close()

	var p []entities.Post
	err = db.Model(&p).
		ColumnExpr("post.*").
		Join("LEFT JOIN likes AS l ON l.post_id = post.id").
		Where("l.user_id = ?", user_id).
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return p, err
}
