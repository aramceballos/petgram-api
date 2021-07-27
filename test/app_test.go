package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/aramceballos/petgram-api/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func getToken(app *fiber.App) string {

	type data struct {
		Token string
	}
	type resBody struct {
		Data    data
		Message string
	}

	postBody, _ := json.Marshal(map[string]string{
		"identity": "tester",
		"password": "test",
	})

	req, _ := http.NewRequest(
		"POST",
		"/api/login",
		bytes.NewBuffer(postBody),
	)
	req.Header.Set("Content-Type", "application/json")

	res, _ := app.Test(req, -1)

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("could not read response: %v", err)
	}

	var got resBody
	err = json.Unmarshal(b, &got)
	if err != nil {
		log.Fatalf("could not unmarshall response %v", err)
	}

	return got.Data.Token
}

func TestCategoriesRoute(t *testing.T) {

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "categories route",
			route:         "/api/categories",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "category route",
			route:         "/api/category/1",
			expectedError: false,
			expectedCode:  200,
		},
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(cors.New())

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	api := app.Group("/api")

	routes.AuthRouter(api)
	routes.CategoriesRouter(api)

	token := getToken(app)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("could not read response: %v", err)
		}

		var got map[string]interface{}
		err = json.Unmarshal(b, &got)
		if err != nil {
			log.Fatalf("could not unmarshall response %v", err)
		}
	}
}

func TestUsersRoute(t *testing.T) {

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "user route",
			route:         "/api/user?id=3",
			expectedError: false,
			expectedCode:  200,
		},
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(cors.New())

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	api := app.Group("/api")

	routes.AuthRouter(api)
	routes.UsersRouter(api)

	token := getToken(app)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("could not read response: %v", err)
		}

		var got map[string]interface{}
		err = json.Unmarshal(b, &got)
		if err != nil {
			log.Fatalf("could not unmarshall response %v", err)
		}
	}
}

func TestPostsRoute(t *testing.T) {

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "posts route",
			route:         "/api/posts",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "post route",
			route:         "/api/post/18",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "favorites route",
			route:         "/api/posts/favorites",
			expectedError: false,
			expectedCode:  200,
		},
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	app.Use(cors.New())

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	api := app.Group("/api")

	routes.AuthRouter(api)
	routes.PostsRouter(api)

	token := getToken(app)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("could not read response: %v", err)
		}

		var got map[string]interface{}
		err = json.Unmarshal(b, &got)
		if err != nil {
			log.Fatalf("could not unmarshall response %v", err)
		}
	}
}
