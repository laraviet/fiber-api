package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	Id     int    `gorm:"primarykey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	app := fiber.New()

	dsn := "root:root@tcp(127.0.0.1:3306)/fiber-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Không thể kết nối tới database")
	}

	db.AutoMigrate(&Book{})

	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Hello World",
		})
	})

	// GET /books
	app.Get("/books", func(c *fiber.Ctx) error {
		var books []Book
		db.Find(&books) // Lấy danh sách books từ database
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": 200,
			"data":   books,
		})
	})

	// POST /books
	app.Post("/books", func(c *fiber.Ctx) error {
		var book Book
		c.BodyParser(&book) // parse request body vào trong model
		db.Create(&book)    // Tạo mới record trong database từ model
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Đã tạo mới book thành công",
		})
	})

	// PUT /books/:id
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		var book Book
		c.BodyParser(&book)                   // parse request body vào trong model
		id, _ := strconv.Atoi(c.Params("id")) // Lấy id từ query tring trên request và convert sang int rồi gán vào id
		db.Model(&Book{}).Where("id = ?", id).Updates(map[string]interface{}{"title": book.Title, "author": book.Author})
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Đã cập nhật book thành công",
		})
	})

	// DELETE /books/:id
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id")) // Lấy id từ query tring trên request và convert sang int rồi gán vào id
		db.Delete(&Book{}, id)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Đã xóa book thành công",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
