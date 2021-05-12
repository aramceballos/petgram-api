package main

import (
	"fmt"

	"github.com/aramceballos/petgram-api/api/routes"
	"github.com/aramceballos/petgram-api/pkg/auth"
	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/aramceballos/petgram-api/pkg/posts"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	postsRepo := posts.NewPostgresRepository()
	postsService := posts.NewService(postsRepo)

	categoriesRepo := categories.NewPostgresRepository()
	categoriesService := categories.NewService(categoriesRepo)

	authRepo := auth.NewPostgresRepository()
	authService := auth.NewService(authRepo)

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
	routes.AuthRouter(api, authService)
	_ = app.Listen(":5000")

}
