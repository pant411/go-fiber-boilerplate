package app

import (
	"fmt"
	config "go-fiber-boilerplate/config/database"
	"go-fiber-boilerplate/helpers/errorHandler"

	"go-fiber-boilerplate/middlewares/response"
	routeAuth "go-fiber-boilerplate/modules/auth/route"
	modelTodo "go-fiber-boilerplate/modules/todo/model"
	routeTodo "go-fiber-boilerplate/modules/todo/route"
	modelUser "go-fiber-boilerplate/modules/user/model"
	routeUser "go-fiber-boilerplate/modules/user/route"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	db.AutoMigrate(&modelTodo.Todo{}, &modelUser.User{})

	// Setup routes
	routeTodo.TodoRoutes(app, db)
	routeAuth.AuthRoutes(app, db)
	routeUser.UserRoutes(app, db)

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		return err
	}

	return nil
}
