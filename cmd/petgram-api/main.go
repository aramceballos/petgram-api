package main

import (
	"fmt"
	"os"

	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var postsController = posts.NewPostsController()
var categoriesController = categories.NewCategoriesController()

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	app.Get("/p", postsController.GetAll)
	app.Get("/p/:id", postsController.GetOne)

	app.Get("/c", categoriesController.GetAll)
	app.Get("/c/:id", categoriesController.GetOne)

	app.Listen(":" + port)
}
