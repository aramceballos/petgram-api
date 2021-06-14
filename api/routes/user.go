package routes

import (
	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app fiber.Router, service users.Service) {
	app.Get("/u", middleware.Protected(), getUser(service))
}

func getUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		username := c.Query("username")
		if username == "" {
			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": "username was not provided",
				"data":    nil,
			})
		}

		user, err := service.FetchUser(username)
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
