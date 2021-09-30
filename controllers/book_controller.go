package controllers

import (
	"fiber-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookController struct {
	Db *gorm.DB
}

func (controller *BookController) Index(c *fiber.Ctx) error {
	var books []models.Book
	controller.Db.Find(&books) // Lấy danh sách books từ database
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   books,
	})
}

func (controller *BookController) Store(c *fiber.Ctx) error {
	var book models.Book
	c.BodyParser(&book)         // parse request body vào trong model
	controller.Db.Create(&book) // Tạo mới record trong database từ model
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Đã tạo mới book thành công",
	})
}

func (controller *BookController) Update(c *fiber.Ctx) error {
	var book models.Book
	c.BodyParser(&book)                   // parse request body vào trong model
	id, _ := strconv.Atoi(c.Params("id")) // Lấy id từ query tring trên request và convert sang int rồi gán vào id
	controller.Db.Model(&models.Book{}).Where("id = ?", id).Updates(map[string]interface{}{"title": book.Title, "author": book.Author})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Đã cập nhật book thành công",
	})
}

func (controller *BookController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) // Lấy id từ query tring trên request và convert sang int rồi gán vào id
	controller.Db.Delete(&models.Book{}, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"message": "Đã xóa book thành công",
	})
}
