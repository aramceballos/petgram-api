package routes

import (
	"fmt"
	"strconv"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/gofiber/fiber/v2"
)

func CategoriesRouter(app fiber.Router, service categories.Service) {
	app.Get("/categories", middleware.Protected(), getCategories(service))
	app.Get("/category/:id", middleware.Protected(), getCategory(service))
}

func getCategories(service categories.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := service.FetchCategories()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    categories,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
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
				"status":  "error",
				"message": err,
				"data":    category,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Posts retrieved",
			"data":    category,
		})
	}
}
