package main

import (
	"go-fiber-api-1/controllers/bookController"
	"go-fiber-api-1/controllers/userController"
	"go-fiber-api-1/model"
	"go-fiber-api-1/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	model.ConnectDatabase()

	app := fiber.New()

	// example: /api/books/id 

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookController.GetAllBooks)
	book.Get("/:id", bookController.GetBookByID)

	book.Post("/", utils.AuthMiddleware, bookController.Create)
	book.Put("/:id", utils.AuthMiddleware, bookController.Update)
	book.Delete("/:id", utils.AuthMiddleware, bookController.Delete)

	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)

	app.Listen(":8000")
}