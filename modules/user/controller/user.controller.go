package controller

import (
	"go-fiber-boilerplate/modules/user/model"
	"go-fiber-boilerplate/modules/user/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *service.UserService
}

// NewTodoController creates a new TodoController instance
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) FindAllUser(ctx *fiber.Ctx) error {
	users, err := c.userService.FindAllUser()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(users)
}

// FindOneUser returns a todo by ID
func (c *UserController) FindOneUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user, err := c.userService.FindOneUser(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user := new(model.User)

	if err := ctx.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := user.Validate(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	user.ID = uint(id)
	if err := c.userService.UpdateUser(user.ID, user); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(user)
}
