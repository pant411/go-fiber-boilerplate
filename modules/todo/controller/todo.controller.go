package controller

import (
	"go-fiber-boilerplate/modules/todo/model"
	"go-fiber-boilerplate/modules/todo/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	todoService *service.TodoService
}

// NewTodoController creates a new TodoController instance
func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (c *TodoController) GetAllTodos(ctx *fiber.Ctx) error {
	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(todos)
}

// CreateTodo creates a new todo
func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	todo := new(model.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := todo.Validate(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	if err := c.todoService.CreateTodo(todo); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(todo)
}

// GetTodoByID returns a todo by ID
func (c *TodoController) GetTodoByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	todo, err := c.todoService.GetTodoByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(todo)
}

// UpdateTodo updates a todo
func (c *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	todo := new(model.Todo)

	if err := ctx.BodyParser(todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := todo.Validate(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	todo.ID = uint(id)
	if err := c.todoService.UpdateTodo(todo.ID, todo); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.JSON(todo)
}

// DeleteTodo deletes a todo
func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := c.todoService.DeleteTodo(uint(id)); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return nil
}
