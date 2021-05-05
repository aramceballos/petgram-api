package categories

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var categoriesService = NewCategoriesService()

type CategoriesController interface {
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
}

type controller struct{}

func NewCategoriesController() CategoriesController {
	return &controller{}
}

func (*controller) GetAll(c *fiber.Ctx) error {
	categories, err := categoriesService.FindAll()

	if err != nil {
		return c.JSON(fiber.Map{
			"data":    nil,
			"message": "Error fetching categories",
		})
	}

	return c.JSON(fiber.Map{
		"data":    categories,
		"message": "Categories retrieved",
	})
}

func (*controller) GetOne(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println("Error casting id to int")
	}

	category, err := categoriesService.Find(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"data":    nil,
			"message": fmt.Sprintf("Category with the id:%d not found", id),
		})
	}

	return c.JSON(fiber.Map{
		"data":    category,
		"message": "Category retrieved",
	})
}
