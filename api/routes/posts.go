package routes

import (
	"net/http"
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
		return nil, false
	}
}

func PostsRouter(app fiber.Router) {
	postsService := posts.NewService()

	app.Get("/posts", middleware.Protected(), getPosts(postsService))
	app.Get("/post/:id", middleware.Protected(), getPost(postsService))
	app.Post("/like", middleware.Protected(), likePost(postsService))
	app.Post("/unlike", middleware.Protected(), unlikePost(postsService))
	app.Get("/posts/favorites", middleware.Protected(), getLikedPosts(postsService))
}

func getPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := strconv.Atoi(c.Query("user_id"))

		if err == nil {
			posts, err := service.FetchPostsByUserID(userId)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
					"status":  "error",
					"message": err.Error(),
					"data":    nil,
				})
			}
			return c.JSON(&fiber.Map{
				"status":  "success",
				"message": "Posts retrieved",
				"data":    posts,
			})
		}

		posts, err := service.FetchPosts()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
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
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": "invalid id",
				"data":    nil,
			})
		}

		post, err := service.FetchPost(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
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

		var userId int = int(claims["sub"].(float64))

		postId, err := strconv.Atoi(c.Query("post_id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": "invalid post_id",
				"data":    nil,
			})
		}

		err = service.LikePost(userId, postId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.Status(http.StatusCreated).JSON(&fiber.Map{
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

		var userId int = int(claims["sub"].(float64))

		post_id, err := strconv.Atoi(c.Query("post_id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": "invalid post_id",
				"data":    nil,
			})
		}

		err = service.UnlikePost(userId, post_id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.Status(http.StatusCreated).JSON(&fiber.Map{
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

		var userId int = int(claims["sub"].(float64))

		likedPosts, err := service.FetchLikedPosts(userId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
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
