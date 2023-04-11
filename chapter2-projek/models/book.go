package models

import "time"

type Book struct {
	ID	   int 			`gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title  string		`gorm:"not null;unique;varchar(50)" json:"title"`
	Author string 		`gorm:"not null;varchar(50)" json:"author"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type BookCreated struct {
	Title  string		`json:"title"`
	Author string 		`json:"author"`
}

type BookUpdated struct {
	Title  string		`json:"title"`
	Author string 		`json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
}