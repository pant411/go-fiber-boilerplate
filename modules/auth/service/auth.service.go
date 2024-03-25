package service

import (
	modelLogin "go-fiber-boilerplate/modules/auth/model"
	"go-fiber-boilerplate/modules/user/model"
	"go-fiber-boilerplate/modules/user/repository"

	"github.com/gofiber/fiber/v2"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// CreateUser creates a new user
func (s *AuthService) Register(user *model.User) error {
	_, err := s.userRepo.FindOneUserByEmail(user.Email)
	if err == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user exist")
	}
	if err := user.HashPassword(); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(credentials *modelLogin.Login) (interface{}, error) {
	user, err := s.userRepo.FindOneUserByEmail(credentials.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "user not exist")
	}
	if err := user.ComparePassword(credentials.Password); err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "wrong password")
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"Username": user.Username,
		"Email":    user.Email,
		"Name":     user.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError)
	}

	return fiber.Map{"accessToken": t}, nil
}
