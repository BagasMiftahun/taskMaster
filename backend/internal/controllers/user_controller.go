package controllers

import (
	"task-master/internal/models"
	logger "task-master/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx, db *gorm.DB, log logger.Logger) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Error("Cannot parse JSON: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := db.Create(&user).Error; err != nil {
		log.Error("Failed to create user: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	log.Info("User created successfully: " + user.Name)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUser(c *fiber.Ctx, db *gorm.DB, log logger.Logger) error {
	id := c.Params("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		log.Error("User not found: " + err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	log.Info("User retrieved successfully: " + user.Name)
	return c.JSON(user)
}

// Implement other controller functions (UpdateUser, DeleteUser)
