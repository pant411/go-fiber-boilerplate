package controller

import (
	"strconv"

	"go-fiber-boilerplate/modules/todo/model"
	"go-fiber-boilerplate/modules/todo/service"

	"github.com/gofiber/fiber/v2"
)

// TodoController handles HTTP requests related to Todo operations
type TodoController struct {
	todoService *service.TodoService
}

// NewTodoController creates a new TodoController instance
func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

// @Summary Get all todos
// @Description Get all todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} Todo
// @Router /api/todo [get]
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	todo := new(model.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	todo.ID = uint(id)
	if err := c.todoService.UpdateTodo(todo); err != nil {
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
	return ctx.JSON(fiber.Map{"message": "Todo deleted successfully"})
}