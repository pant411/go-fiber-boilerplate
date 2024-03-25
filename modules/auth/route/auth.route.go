package routes

import (
	"go-fiber-boilerplate/modules/auth/controller"
	"go-fiber-boilerplate/modules/auth/service"
	"go-fiber-boilerplate/modules/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	// Define routes for Todo module
	authRoutes := app.Group("/api/auth")

	authRoutes.Post("/register", authController.Register)
	authRoutes.Post("/login", authController.Login)
}
