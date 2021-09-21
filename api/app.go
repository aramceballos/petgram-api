package main

import (
	"log"
	"os"

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
		log.Fatal("Error loading .env file")
	}

	api := app.Group("/api")

	routes.PostsRouter(api)
	routes.CategoriesRouter(api)
	routes.AuthRouter(api)
	routes.UsersRouter(api)

	port := ":" + os.Getenv("PORT")

	app.Listen(port)

}
