package main

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var postsController = posts.NewPostsController()
var categoriesController = categories.NewCategoriesController()

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/p", func(c *fiber.Ctx) error {
		return c.SendString("Not Found")
	})

	app.Get("/p/:id", postsController.GetOne)
	app.Get("/c/:id", categoriesController.GetOne)

	app.Listen(":" + port)
}
