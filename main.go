package main

import (
	"fiber-api/databases"
	"fiber-api/routes"
	"log"
)

func main() {
	db := databases.SetupDB()
	databases.Migrate(db)

	app := routes.SetupRoutes(db)

	log.Fatal(app.Listen(":3000"))
}
