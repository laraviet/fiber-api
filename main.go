package main

import (
	"fiber-api/databases"
	"fiber-api/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := databases.SetupDB()
	databases.Migrate(db)

	app := routes.SetupRoutes(db)

	log.Fatal(app.Listen(":3000"))
}
