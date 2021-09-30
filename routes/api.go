package routes

import (
	"fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	bookController := controllers.BookController{
		Db: db,
	}

	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Hello World",
		})
	})

	app.Get("/books", bookController.Index)         // GET /books
	app.Post("/books", bookController.Store)        // POST /books
	app.Put("/books/:id", bookController.Update)    // PUT /books/:id
	app.Delete("/books/:id", bookController.Delete) // DELETE /books/:id

	return app
}
