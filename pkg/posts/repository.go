package posts

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/go-pg/pg/v10"
)

type Repository interface {
	ReadPosts() ([]entities.Post, error)
	ReadPostsByUserID(int) ([]entities.Post, error)
	ReadPost(int) (entities.Post, error)
	LikePost(int, int) error
	UnlikePost(int, int) error
	ReadLikedPosts(int) ([]entities.Post, error)
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

func (r *repo) ReadPosts() ([]entities.Post, error) {
	opt, err := pg.ParseURL(r.url)
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

func (r *repo) ReadPostsByUserID(userId int) ([]entities.Post, error) {
	opt, err := pg.ParseURL(r.url)
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
		Where("post.user_id = ?", userId).
		Relation("Likes").
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return p, err
}

func (r *repo) ReadPost(id int) (entities.Post, error) {
	opt, err := pg.ParseURL(r.url)
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

func (r *repo) LikePost(userId int, postId int) error {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return err
	}

	db := pg.Connect(opt)
	defer db.Close()

	l := &entities.Like{UserID: userId, PostID: postId}
	_, err = db.Model(l).Insert()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (r *repo) UnlikePost(userId int, postId int) error {
	opt, err := pg.ParseURL(r.url)
	if err != nil {
		fmt.Println("Unable to connect to database")
		return err
	}

	db := pg.Connect(opt)
	defer db.Close()

	l := &entities.Like{}
	_, err = db.Model(l).Where("user_id = ? AND post_id = ?", userId, postId).Delete()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (r *repo) ReadLikedPosts(userId int) ([]entities.Post, error) {
	opt, err := pg.ParseURL(r.url)
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
		Where("l.user_id = ?", userId).
		Select()
	if err != nil {
		fmt.Println(err)
	}

	return p, err
}
