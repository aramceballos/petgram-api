package routes

import (
	"net/http"
	"strconv"

	"github.com/aramceballos/petgram-api/api/middleware"
	"github.com/aramceballos/petgram-api/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app fiber.Router) {
	usersService := users.NewService()

	app.Get("/user", middleware.Protected(), getUser(usersService))
}

func getUser(service users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		username := c.Query("username")
		userId, err := strconv.Atoi(c.Query("id"))

		if username != "" {
			user, err := service.FetchUser(username)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
					"status":  "error",
					"message": err.Error(),
					"data":    nil,
				})
			}

			return c.JSON(&fiber.Map{
				"status":  "success",
				"message": "User retrieved",
				"data":    user,
			})
		}

		if err == nil {
			user, err := service.FetchUserById(userId)
			if err != nil {
				return c.JSON(&fiber.Map{
					"status":  "error",
					"message": err.Error(),
					"data":    nil,
				})
			}

			return c.JSON(&fiber.Map{
				"status":  "success",
				"message": "User retrieved",
				"data":    user,
			})
		}

		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "username or user_id were not provided not provided",
			"data":    nil,
		})
	}
}
