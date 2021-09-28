# Installation Steps
1. Tạo 1 folder mới `fiber-api`
2. cd vào folder mới trên `cd fiber-api`
3. Chạy các command dưới đây
- go mod init fiber-api
- go get -u github.com/gofiber/fiber/v2
- go mod download
- go mod vendor
4. Tạo file .gitignore với nội dung dưới đây
```
vendor/
```
5. Tạo file `main.go`
```
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Hello World",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
```
6. Ở root folder chạy command sau
- `go run main.go`
7. Mở `postman` và tạo GET request tới `http://localhost:3000` và kiểm trả response