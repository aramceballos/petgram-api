package posts

import (
	"database/sql"
	"log"
	"os"

	"github.com/aramceballos/petgram-api/pkg/entities"
	_ "github.com/lib/pq"
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

func (r *repo) ReadPosts() ([]entities.Post, error) {
	var posts []entities.Post

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT posts.*, u.name, u.username, u.email FROM posts JOIN users AS u ON u.id = posts.user_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.PostDate, &post.ImageURL, &post.Description, &post.Name, &post.Username, &post.Email)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	rows, err = tx.Query("SELECT * from likes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var like entities.Like
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID)
		if err != nil {
			return nil, err
		}
		for i, p := range posts {
			if p.ID == like.PostID {
				posts[i].Likes = append(p.Likes, &like)
			}
		}
	}

	err = tx.Rollback()

	return posts, err
}

func (r *repo) ReadPostsByUserID(userId int) ([]entities.Post, error) {
	var posts []entities.Post

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT posts.*, u.name, u.username, u.email FROM posts JOIN users AS u ON u.id = posts.user_id WHERE posts.user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.PostDate, &post.ImageURL, &post.Description, &post.Name, &post.Username, &post.Email)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	rows, err = tx.Query("SELECT * from likes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var like entities.Like
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID)
		if err != nil {
			return nil, err
		}

		for i, p := range posts {
			if p.ID == like.PostID {
				posts[i].Likes = append(posts[i].Likes, &like)
			}
		}
	}

	err = tx.Rollback()

	return posts, err
}

func (r *repo) ReadPost(id int) (entities.Post, error) {
	post := &entities.Post{ID: id}

	tx, err := r.db.Begin()
	if err != nil {
		return entities.Post{}, err
	}

	err = tx.QueryRow("SELECT posts.*, u.name, u.username, u.email FROM posts JOIN users AS u ON u.id = posts.user_id WHERE posts.id = $1", id).Scan(&post.ID, &post.UserID, &post.CategoryID, &post.PostDate, &post.ImageURL, &post.Description, &post.Name, &post.Username, &post.Email)
	if err != nil {
		return entities.Post{}, err
	}

	rows, err := tx.Query("SELECT * from likes")
	if err != nil {
		return entities.Post{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var like entities.Like
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID)
		if err != nil {
			return entities.Post{}, err
		}

		if post.ID == like.PostID {
			post.Likes = append(post.Likes, &like)
		}
	}

	return *post, err
}

func (r *repo) LikePost(userId int, postId int) error {
	stmt, err := r.db.Prepare("INSERT INTO likes (user_id, post_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, postId)

	return err
}

func (r *repo) UnlikePost(userId int, postId int) error {
	stmt, err := r.db.Prepare("DELETE FROM likes WHERE user_id = $1 AND post_id = $2")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, postId)

	return err
}

func (r *repo) ReadLikedPosts(userId int) ([]entities.Post, error) {
	var posts []entities.Post

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT posts.*, u.name, u.username, u.email FROM posts JOIN users AS u ON u.id = posts.user_id JOIN likes ON likes.post_id = posts.id WHERE likes.user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.CategoryID, &post.PostDate, &post.ImageURL, &post.Description, &post.Name, &post.Username, &post.Email)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	rows, err = tx.Query("SELECT * from likes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var like entities.Like
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID)
		if err != nil {
			return nil, err
		}

		for i, p := range posts {
			if p.ID == like.PostID {
				posts[i].Likes = append(posts[i].Likes, &like)
			}
		}
	}

	return posts, err
}
