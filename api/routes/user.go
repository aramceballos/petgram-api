package routes

import (
	"fmt"
	"strconv"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app fiber.Router, service users.Service) {
	app.Get("/u", middleware.Protected(), getUser(service))
}

func getUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			fmt.Println("Error casting id to int")
		}

		user, err := service.FetchUser(id)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "User retrieved",
			"data":    user,
		})
	}
}
