package categories

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var categoriesService = NewCategoriesService()

type CategoriesController interface {
	GetOne(c *fiber.Ctx) error
}

type controller struct{}

func NewCategoriesController() CategoriesController {
	return &controller{}
}

func (*controller) GetOne(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println("Error casting id to int")
	}

	category, err := categoriesService.Find(id)

	return c.JSON(fiber.Map{
		"data":    category,
		"message": "Category retrieved",
	})
}
