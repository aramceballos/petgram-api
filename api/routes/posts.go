package routes

import (
	"fmt"
	"strconv"

	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func PostsRouter(app fiber.Router, service posts.Service) {
	app.Get("/p", getPosts(service))
	app.Get("/p/:id", getPost(service))
}

func CategoriesRouter(app fiber.Router, service categories.Service) {
	app.Get("/c", getCategories(service))
	app.Get("/c/:id", getCategory(service))
}

func UsersRouter(app fiber.Router, service users.Service) {
	app.Post("/signup", createUser(service))
}

func getPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := service.FetchPosts()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"data":    posts,
				"message": err,
			})
		}
		return c.JSON(&fiber.Map{
			"data":    posts,
			"message": "Posts retrieved",
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
				"posts":   posts,
				"message": err,
			})
		}
		return c.JSON(&fiber.Map{
			"posts":   posts,
			"message": "Posts retrieved",
		})
	}
}

func getCategories(service categories.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := service.FetchCategories()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"data":    categories,
				"message": err,
			})
		}
		return c.JSON(&fiber.Map{
			"data":    categories,
			"message": "Posts retrieved",
		})
	}
}

func getCategory(service categories.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			fmt.Println("Error casting id to int")
		}

		category, err := service.FetchCategory(id)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"data":    category,
				"message": err,
			})
		}
		return c.JSON(&fiber.Map{
			"data":    category,
			"message": "Posts retrieved",
		})
	}
}

func createUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(entities.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		err := service.InsertUser(user)

		if err != nil {
			return c.JSON(fiber.Map{
				"data":    nil,
				"message": "Error creating user",
			})
		}

		return c.JSON(fiber.Map{
			"data":    nil,
			"message": "User created",
		})
	}
}
