package auth

import (
	"errors"
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

	var ud entities.User

	identity := input.Identity
	pass := input.Password

	email, err := s.repository.ReadUserByEmail(identity)
	if err == nil {
		ud = entities.User{
			ID:       email.ID,
			Name:     email.Name,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	} else {
		user, err := s.repository.ReadUserByUsername(identity)
		if err == nil {
			ud = entities.User{
				ID:       user.ID,
				Name:     user.Name,
				Username: user.Username,
				Email:    user.Email,
				Password: user.Password,
			}
		} else {
			return entities.Response{}, errors.New("user not found")
		}
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return entities.Response{}, errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secret := os.Getenv("SECRET")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return entities.Response{}, errors.New("error signing token")
	}

	res := entities.Response{
		ID:       ud.ID,
		Name:     ud.Name,
		Username: ud.Username,
		Token:    t,
	}

	return res, err
}

func (s *service) InsertUser(user *entities.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Unable to has password")
	}
	user.Password = string(hashedPass)

	e := user.Email
	u := user.Username

	email, _ := s.repository.ReadUserByEmail(e)
	if email != nil {
		return errors.New("user already exists")
	}

	userName, _ := s.repository.ReadUserByUsername(u)
	if userName != nil {
		return errors.New("user already exists")
	}

	err = s.repository.CreateUser(user)

	return err
}
