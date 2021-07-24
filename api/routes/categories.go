package routes

import (
	"net/http"
	"strconv"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/categories"
	"github.com/gofiber/fiber/v2"
)

func CategoriesRouter(app fiber.Router) {
	categoriesService := categories.NewService()

	app.Get("/categories", middleware.Protected(), getCategories(categoriesService))
	app.Get("/category/:id", middleware.Protected(), getCategory(categoriesService))
}

func getCategories(service categories.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := service.FetchCategories()
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Categories retrieved",
			"data":    categories,
		})
	}
}

func getCategory(service categories.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": "invalid id",
				"data":    nil,
			})
		}

		category, err := service.FetchCategory(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Category retrieved",
			"data":    category,
		})
	}
}
