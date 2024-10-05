package app

import (
	"task-master/config"
	"task-master/internal/routes"
	logger "task-master/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewApp(log logger.Logger) *fiber.App {
	app := fiber.New()

	// Initialize GORM
	dsn := config.AppConfig.Database.User + ":" +
		config.AppConfig.Database.Password + "@tcp(" +
		config.AppConfig.Database.Host + ":" +
		config.AppConfig.Database.Port + ")/" +
		config.AppConfig.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error()) // Perbaikan di sini
	}

	// Setup middleware
	app.Use(func(c *fiber.Ctx) error {
		log.Info("Request: " + c.Method() + " " + c.Path())
		return c.Next()
	})

	// Setup routes with DB
	routes.SetupRoutes(app, db, log)

	return app
}
