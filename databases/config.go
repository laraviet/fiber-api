package databases

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	connection := os.Getenv("DB_CONNECTION")
	var db *gorm.DB
	var err error

	if connection == "mysql" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})	
	}
	
	if err != nil {
		log.Fatal("Không thể kết nối tới database")
	}
	return db
}
