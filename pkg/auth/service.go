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
	ReadUser(entities.LoginInput) (string, error)
	InsertUser(*entities.User) error
}

type service struct {
	repository Repository
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) ReadUser(input entities.LoginInput) (string, error) {

	var ud entities.User

	identity := input.Identity
	pass := input.Password

	email, err := s.repository.ReadUserByEmail(identity)
	if err != nil {
		return "", errors.New("error on email")
	}

	user, err := s.repository.ReadUserByUsername(identity)
	if err != nil {
		return "", errors.New("error on username")
	}

	if email == nil && user == nil {
		return "", errors.New("user not found")
	}

	if email == nil {
		ud = entities.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = entities.User{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !CheckPasswordHash(pass, ud.Password) {
		return "", errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secret := os.Getenv("SECRET")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("error signing token")
	}

	return t, err
}

func (s *service) InsertUser(user *entities.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Unable to has password")
	}
	user.Password = string(hashedPass)

	err = s.repository.CreateUser(user)

	return err
}
