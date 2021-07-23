package main

import (
	"fmt"

	"github.com/aramceballos/petgram-api/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(cors.New())

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	api := app.Group("/api")

	routes.PostsRouter(api)
	routes.CategoriesRouter(api)
	routes.AuthRouter(api)
	routes.UsersRouter(api)

	app.Listen(":5000")

}
