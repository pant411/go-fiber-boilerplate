package cmd

import (
	"fmt"
	config "go-fiber-boilerplate/config/database"
	"go-fiber-boilerplate/helpers/errorHandler"

	"go-fiber-boilerplate/middlewares/response"
	"go-fiber-boilerplate/modules/todo/model"
	routes "go-fiber-boilerplate/modules/todo/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

func Execute() error {
	// Load configuration
	cfgDB, err := config.LoadConfigDB()
	if err != nil {
		return err
	}

	var PORT string = os.Getenv("PORT")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: errorHandler.AllException,
	})

	app.Use(response.ModifyJSONResponse())

	// Connect to MySQL database
	db, err := gorm.Open(mysql.Open(cfgDB.GetDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the database
	db.AutoMigrate(&model.Todo{})

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Setup routes
	routes.TodoRoutes(app, db)

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		return err
	}

	return nil
}
