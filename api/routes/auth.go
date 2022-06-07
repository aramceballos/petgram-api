package routes

import (
	"net/http"

	"github.com/aramceballos/petgram-api/pkg/auth"
	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router) {
	authService := auth.NewService()

	app.Post("/login", login(authService))
	app.Post("/signup", signup(authService))
}

func login(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "invalid body",
				"data":    nil,
			})
		}

		res, err := service.ReadUser(input)
		if err != nil {
			if err.Error() == "error signing token" {
				return c.SendStatus(http.StatusInternalServerError)
			}

			return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "Success login",
			"data":    res,
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
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"message": "User created",
			"data":    nil,
		})
	}
}
