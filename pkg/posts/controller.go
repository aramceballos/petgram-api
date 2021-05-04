package posts

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var postsService = NewPostsService()

type PostsController interface {
	GetPost(c *fiber.Ctx) error
}

type controller struct{}

func NewPostsController() PostsController {
	return &controller{}
}

func (*controller) GetPost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println("Error casting id to int")
	}

	post, err := postsService.Find(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"data":    nil,
			"message": fmt.Sprintf("Post with the id:%d not found", id),
		})
	}

	return c.JSON(fiber.Map{
		"data":    post,
		"message": "Post retrieved",
	})
}
