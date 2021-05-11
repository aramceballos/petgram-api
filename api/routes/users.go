package routes

import (
	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app fiber.Router, service users.Service) {
	app.Post("/signup", createUser(service))
}

func createUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(entities.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		err := service.InsertUser(user)

		if err != nil {
			return c.JSON(fiber.Map{
				"data":    nil,
				"message": "Error creating user",
			})
		}

		return c.JSON(fiber.Map{
			"data":    nil,
			"message": "User created",
		})
	}
}
