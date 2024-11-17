package userController

import (
	"go-fiber-api-1/model"
	"go-fiber-api-1/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user model.User

	// Parsing request body ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Cek apakah email sudah terdaftar
	var existingUser model.User
	if err := model.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already in use",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user.Password = hashedPassword // Simpan password yang sudah di-hash

	// Simpan user ke database
	if err := model.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":    user.Id,
			"email": user.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var user model.User

	// Parsing request body ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Cek apakah email ada di database
	var dbUser model.User
	if err := model.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Verifikasi password
	if err := utils.CheckPasswordHash(user.Password, dbUser.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate token JWT
	token, err := utils.GenerateToken(dbUser.Email, dbUser.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
