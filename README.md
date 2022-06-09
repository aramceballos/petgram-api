# Petgram API

API for Petgram

## Usage

### Configuration

In order to connect to postgres you need to set some environment variables that are listed in .env.example

### Installation

To run the server you can use docker

```shell
$ docker build -t petgram-api .
$ docker run -p 5000:5000 --name petgram-api petgram-api
```

Or the terminal

```shell
$ go build api/app.go
```

## Endpoints

| Method | URL                 | Protected | Action                                                            |
| ------ | ------------------- | --------- | ----------------------------------------------------------------- |
| GET    | "/categories"       | Yes       | Returns a list of categories                                      |
| GET    | "/category/:id"     | Yes       | Returns a category with the given id                              |
| GET    | "/posts"            | Yes       | Returns a list of posts                                           |
| GET    | "/posts?user_id=id" | Yes       | Returns a list of posts posted by the user with the given user_id |
| GET    | "/post/:id"         | Yes       | Returns a post with the given id                                  |
| POST   | "/like"             | Yes       | Likes a post                                                      |
| POST   | "/unlike"           | Yes       | Unlikes a post                                                    |
| GET    | "/posts/favorites"  | Yes       | Returns a list of liked posts                                     |
| GET    | "/user"             | Yes       | Returns a user with the username or id passed by query strings    |
| POST   | "/login"            | No        | Returns an access token                                           |
| POST   | "/signup"           | No        | Creates a new user                                                |
