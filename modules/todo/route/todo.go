package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/controller"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/repository"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/service"
	"gorm.io/gorm"
)

// SetupTodoRoutes sets up routes for the Todo module
func SetupTodoRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository, service, and controller
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	// Define routes for Todo module
	todoRoutes := app.Group("/api/todo")
	todoRoutes.Get("/", todoController.GetAllTodos)
	todoRoutes.Post("/", todoController.CreateTodo)
	todoRoutes.Get("/:id", todoController.GetTodoByID)
	todoRoutes.Patch("/:id", todoController.UpdateTodo)
	todoRoutes.Delete("/:id", todoController.DeleteTodo)

	// Add more routes here as needed
}