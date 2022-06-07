package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/aramceballos/petgram-api/pkg/entities"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	ReadUser(entities.LoginInput) (entities.Response, error)
	InsertUser(*entities.User) error
}

type service struct {
	repository Repository
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewService() Service {
	postgresRepository := NewPostgresRepository()

	return &service{
		repository: postgresRepository,
	}
}

func (s *service) ReadUser(input entities.LoginInput) (entities.Response, error) {
	var userData entities.User

	email, err := s.repository.ReadUserByEmail(input.Identity)
	if err == nil {
		userData = entities.User{
			ID:       email.ID,
			Name:     email.Name,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	} else {
		user, err := s.repository.ReadUserByUsername(input.Identity)
		if err == nil {
			userData = entities.User{
				ID:       user.ID,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Password: user.Password,
			}
		} else {
			return entities.Response{}, fmt.Errorf("user not found")
		}
	}

	if !CheckPasswordHash(input.Password, userData.Password) {
		return entities.Response{}, fmt.Errorf("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["sub"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secret := os.Getenv("SECRET")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return entities.Response{}, fmt.Errorf("error signing token")
	}

	res := entities.Response{
		ID:       userData.ID,
		Name:     userData.Name,
		Username: userData.Username,
		Token:    t,
	}

	return res, err
}

func (s *service) InsertUser(user *entities.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Unable to hash password")
		return fmt.Errorf("unable to hash password")
	}
	user.Password = string(hashedPass)

	e := user.Email
	u := user.Username

	_, err = s.repository.ReadUserByEmail(e)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	_, err = s.repository.ReadUserByUsername(u)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	err = s.repository.CreateUser(user)

	return err
}
