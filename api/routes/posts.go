package routes

import (
	"fmt"
	"strconv"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/gofiber/fiber/v2"
)

func PostsRouter(app fiber.Router, service posts.Service) {
	app.Get("/p", middleware.Protected(), getPosts(service))
	app.Get("/p/:id", middleware.Protected(), getPost(service))
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

		posts, err := service.FetchPost(id)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"posts":   posts,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Posts retrieved",
			"posts":   posts,
		})
	}
}
