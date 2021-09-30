package models

type Book struct {
	Id     int    `gorm:"primarykey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
