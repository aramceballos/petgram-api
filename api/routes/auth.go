package routes

import (
	"github.com/aramceballos/petgram-api/pkg/auth"
	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Post("/login", login(service))
	app.Post("/signup", signup(service))
}

func login(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Error on login request",
				"data":    nil,
			})
		}

		t, err := service.ReadUser(input)
		if err != nil {
			if err.Error() == "error on email" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Error on email",
					"data":    nil,
				})
			}

			if err.Error() == "error on username" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Error on username",
					"data":    nil,
				})
			}

			if err.Error() == "user not found" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "User not found",
					"data":    nil,
				})
			}

			if err.Error() == "invalid password" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid password",
					"data":    nil,
				})
			}

			if err.Error() == "error signing token" {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			return c.JSON(&fiber.Map{
				"status":  "error",
				"message": err,
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Success login",
			"data":    t,
		})
	}
}

func signup(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(entities.User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		err := service.InsertUser(user)

		if err != nil {
			return c.JSON(fiber.Map{
				"status":  "error",
				"message": "Error creating user",
				"data":    nil,
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User created",
			"data":    nil,
		})
	}
}
