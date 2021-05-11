package main

import (
	"fmt"

	"github.com/aramceballos/petgram-api/api/routes"
	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	postsRepo := posts.NewPostgresRepository()
	categoriesRepo := categories.NewPostgresRepository()
	usersRepo := users.NewPostgresRepository()
	postsService := posts.NewService(postsRepo)
	categoriesService := categories.NewService(categoriesRepo)
	usersService := users.NewService(usersRepo)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api := app.Group("/api")
	routes.PostsRouter(api, postsService)
	routes.CategoriesRouter(api, categoriesService)
	routes.UsersRouter(api, usersService)
	_ = app.Listen(":5000")

}
