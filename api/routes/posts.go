package routes

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	secret := os.Getenv("SECRET")
	hmacSecret := []byte(secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func PostsRouter(app fiber.Router, service posts.Service) {
	app.Get("/p", middleware.Protected(), getPosts(service))
	app.Get("/p/individual/:id", middleware.Protected(), getPost(service))
	app.Post("/p/l", middleware.Protected(), likePost(service))
	app.Post("/p/ul", middleware.Protected(), unlikePost(service))
	app.Get("/p/f", middleware.Protected(), getLikedPosts(service))
}

func getPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := service.FetchPosts()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    posts,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Posts retrieved",
			"data":    posts,
		})
	}
}

func getPost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			fmt.Println("Error casting id to int")
		}

		post, err := service.FetchPost(id)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Posts retrieved",
			"data":    post,
		})
	}
}

func likePost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		token := strings.Split(authorizationHeader, " ")[1]
		claims, _ := extractClaims(token)

		var userId int = int(claims["user_id"].(float64))

		postId, err := strconv.Atoi(c.Query("post_id"))
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": "Error with post_id query string",
				"data":    nil,
			})
		}

		err = service.LikePost(userId, postId)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Post liked",
			"data":    nil,
		})
	}
}

func unlikePost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		token := strings.Split(authorizationHeader, " ")[1]
		claims, _ := extractClaims(token)

		var userId int = int(claims["user_id"].(float64))

		post_id, err := strconv.Atoi(c.Query("post_id"))
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": "Error with post_id query string",
				"data":    nil,
			})
		}

		err = service.UnlikePost(userId, post_id)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Post unliked",
			"data":    nil,
		})
	}
}

func getLikedPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		token := strings.Split(authorizationHeader, " ")[1]
		claims, _ := extractClaims(token)

		var userId int = int(claims["user_id"].(float64))

		likedPosts, err := service.FetchLikedPosts(userId)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    likedPosts,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Favorites retrieved",
			"data":    likedPosts,
		})
	}
}
