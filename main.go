package main

import (
	"go-fiber-api-1/controllers/bookController"
	"go-fiber-api-1/model"

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
	book.Post("/", bookController.Create)
	book.Put("/:id", bookController.Update)
	book.Delete("/:id", bookController.Delete)

	app.Listen(":8000")
}