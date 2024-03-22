package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/model"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/service"
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.JSON(todos)
}

// CreateTodo creates a new todo
func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	todo := new(model.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	if err := c.todoService.CreateTodo(todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create todo"})
	}
	return ctx.JSON(todo)
}

// GetTodoByID returns a todo by ID
func (c *TodoController) GetTodoByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	todo, err := c.todoService.GetTodoByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	todo.ID = uint(id)
	if err := c.todoService.UpdateTodo(todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update todo"})
	}
	return ctx.JSON(todo)
}

// DeleteTodo deletes a todo
func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	if err := c.todoService.DeleteTodo(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete todo"})
	}
	return ctx.JSON(fiber.Map{"message": "Todo deleted successfully"})
}
