package routes

import (
	"go-fiber-boilerplate/modules/todo/controller"
	"go-fiber-boilerplate/modules/todo/repository"
	"go-fiber-boilerplate/modules/todo/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TodoRoutes sets up routes for the Todo module
func TodoRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository, service, and controller
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	// Define routes for Todo module
	todoRoutes := app.Group("/api/todo")

	// @Summary Get all todos
	// @Description Get all todos
	// @Tags Todos
	// @Accept json
	// @Produce json
	// @Success 200 {array} Todo
	// @Router /api/todo [get]
	todoRoutes.Get("/", todoController.GetAllTodos)

	// @Summary Create a new todo
	// @Description Create a new todo
	// @Tags Todos
	// @Accept json
	// @Produce json
	// @Param todo body TodoCreateRequest true "Todo object to be created"
	// @Success 200 {object} Todo
	// @Router /api/todo [post]
	todoRoutes.Post("/", todoController.CreateTodo)

	// @Summary Get a todo by ID
	// @Description Get a todo by ID
	// @Tags Todos
	// @Accept json
	// @Produce json
	// @Param id path string true "Todo ID"
	// @Success 200 {object} Todo
	// @Router /api/todo/{id} [get]
	todoRoutes.Get("/:id", todoController.GetTodoByID)

	// @Summary Update a todo
	// @Description Update a todo
	// @Tags Todos
	// @Accept json
	// @Produce json
	// @Param id path string true "Todo ID"
	// @Param todo body TodoUpdateRequest true "Updated todo object"
	// @Success 200 {object} Todo
	// @Router /api/todo/{id} [patch]
	todoRoutes.Patch("/:id", todoController.UpdateTodo)

	// @Summary Delete a todo
	// @Description Delete a todo
	// @Tags Todos
	// @Accept json
	// @Produce json
	// @Param id path string true "Todo ID"
	// @Success 204
	// @Router /api/todo/{id} [delete]
	todoRoutes.Delete("/:id", todoController.DeleteTodo)

	// Add more routes here as needed
}
