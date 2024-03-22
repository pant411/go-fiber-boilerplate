package cmd

import (
	config "github.com/pant411/go-fiber-boilerplate/config/database"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/pant411/go-fiber-boilerplate/modules/todo/model"
	routes "github.com/pant411/go-fiber-boilerplate/modules/todo/route"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SwaggerConfig = swagger.Config{
	BasePath: "/",
	FilePath: "./docs/swagger.json",
	Path:     "swagger",
	Title:    "Swagger API Docs",
}

func Execute() error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize Fiber app
	app := fiber.New()

	app.Use(swagger.New(SwaggerConfig))

	// Connect to MySQL database
	db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the database
	db.AutoMigrate(&model.Todo{})

	// Setup routes
	routes.SetupTodoRoutes(app, db)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		return err
	}

	return nil
}
