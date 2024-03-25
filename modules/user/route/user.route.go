package routes

import (
	middleware "go-fiber-boilerplate/middlewares/auth"
	"go-fiber-boilerplate/modules/user/controller"
	"go-fiber-boilerplate/modules/user/repository"
	"go-fiber-boilerplate/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Define routes for Todo module
	userRoutes := app.Group("/api/user")

	userRoutes.Get("/", middleware.Protected(), userController.FindAllUser)
	userRoutes.Get("/:id", middleware.Protected(), userController.FindOneUser)
	userRoutes.Patch("/:id", middleware.Protected(), userController.UpdateUser)
}
