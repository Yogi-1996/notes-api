package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"not null; uniquekey"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
