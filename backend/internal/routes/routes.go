package routes

import (
	"task-master/internal/controllers"
	logger "task-master/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, log logger.Logger) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", func(c *fiber.Ctx) error {
		return controllers.CreateUser(c, db, log)
	})
	api.Get("/users/:id", func(c *fiber.Ctx) error {
		return controllers.GetUser(c, db, log)
	})
	// api.Put("/users/:id", func(c *fiber.Ctx) error {
	// 	return controllers.UpdateUser(c, db, log)
	// })
	// api.Delete("/users/:id", func(c *fiber.Ctx) error {
	// 	return controllers.DeleteUser(c, db, log)
	// })

	// Product routes (tambahkan sesuai kebutuhan)
}
