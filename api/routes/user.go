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

		if username != "" { // If username is not empty, get user by username
			user, err := service.FetchUser(username)
			if err != nil {
				if err.Error() == "user not found" {
					return c.Status(http.StatusNotFound).JSON(&fiber.Map{
						"status":  "error",
						"message": err.Error(),
						"data":    nil,
					})
				}

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
		} else if err == nil { // If id is not empty, get user by id
			user, err := service.FetchUserById(userId)
			if err != nil {
				if err.Error() == "user not found" {
					return c.Status(http.StatusNotFound).JSON(&fiber.Map{
						"status":  "error",
						"message": err.Error(),
						"data":    nil,
					})
				}

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

		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "username or id were not provided not provided",
			"data":    nil,
		})
	}
}
