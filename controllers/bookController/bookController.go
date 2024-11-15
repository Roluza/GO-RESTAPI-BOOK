package bookController

import (
	"go-fiber-api-1/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBooks(c *fiber.Ctx) error {
	var books []model.Book
	model.DB.Find(&books)

	return c.JSON(books);
}

func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book
	if err := model.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Book Not Found!",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Book Not Found!",
		})
	}

	return c.JSON(book)
}

func Create(c *fiber.Ctx) error {
	var book model.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := model.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Book created successfully"})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	if model.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot update data book",
		})
	}

	return c.JSON(fiber.Map{"message": "Book update successfully"})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var book model.Book

	if model.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cannot delete data book",
		})
	}

	return c.JSON(fiber.Map{"message": "Book delete successfully"})
}