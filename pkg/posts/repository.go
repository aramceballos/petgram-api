package posts

import (
	"errors"
	"log"

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
	db pg.DB
}

var postgresRepo *repo

func NewPostgresRepository() Repository {
	url := "postgres://postgres:postgress@db:5432/postgres?sslmode=disable"

	if postgresRepo == nil {
		opt, err := pg.ParseURL(url)
		if err != nil {
			log.Fatal("error parsing db url")
		}
		db := pg.Connect(opt)
		postgresRepo = &repo{
			db: *db,
		}
	}

	return postgresRepo
}

func (r *repo) ReadPosts() ([]entities.Post, error) {
	var p []entities.Post
	err := r.db.Model(&p).
		ColumnExpr("post.*").
		ColumnExpr("u.name, u.username, u.email").
		Join("JOIN petgram.users AS u ON u.id = post.user_id").
		Relation("Likes").
		Select()

	return p, err
}

func (r *repo) ReadPostsByUserID(userId int) ([]entities.Post, error) {
	var p []entities.Post
	err := r.db.Model(&p).
		ColumnExpr("post.*").
		ColumnExpr("u.name, u.username, u.email").
		Join("JOIN petgram.users AS u ON u.id = post.user_id").
		Where("post.user_id = ?", userId).
		Relation("Likes").
		Select()

	if len(p) == 0 {
		return []entities.Post{}, errors.New("there are not posts from the given user")
	}

	return p, err
}

func (r *repo) ReadPost(id int) (entities.Post, error) {
	p := &entities.Post{ID: id}
	err := r.db.Model(p).
		ColumnExpr("post.*").
		ColumnExpr("u.name, u.username, u.email").
		WherePK().
		Join("JOIN petgram.users AS u ON u.id = post.user_id").
		Relation("Likes").
		Select()

	return *p, err
}

func (r *repo) LikePost(userId int, postId int) error {
	l := &entities.Like{UserID: userId, PostID: postId}
	_, err := r.db.Model(l).Insert()

	return err
}

func (r *repo) UnlikePost(userId int, postId int) error {
	l := &entities.Like{}
	_, err := r.db.Model(l).Where("user_id = ? AND post_id = ?", userId, postId).Delete()

	return err
}

func (r *repo) ReadLikedPosts(userId int) ([]entities.Post, error) {
	var p []entities.Post
	err := r.db.Model(&p).
		ColumnExpr("post.*").
		Join("LEFT JOIN petgram.likes AS l ON l.post_id = post.id").
		Where("l.user_id = ?", userId).
		Select()

	return p, err
}
