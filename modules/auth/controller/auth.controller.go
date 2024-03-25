package controller

import (
	modelLogin "go-fiber-boilerplate/modules/auth/model"
	"go-fiber-boilerplate/modules/auth/service"
	"go-fiber-boilerplate/modules/user/model"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	user := new(model.User)
	if err := ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.Validate(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	if err := c.authService.Register(user); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(user)
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	credentials := new(modelLogin.Login)
	if err := ctx.BodyParser(credentials); err != nil {
		return err
	}
	token, err := c.authService.Login(credentials)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(token)
}
