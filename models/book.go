package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NameBook  string    `json:"name_book" gorm:"not null;type:varchar(191)"`
	Author    string    `json:"author" gorm:"not null;type:varchar(191)"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
