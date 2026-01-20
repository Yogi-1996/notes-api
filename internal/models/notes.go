package models

import "time"

type Note struct {
	ID        int       `json:"id" gorm:"primarykey"`
	UserID    int       `json:"user_id" gorm:"not null; uniquekey"`
	Title     string    `json:"title" gorm:"not null" `
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
