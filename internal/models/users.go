package models

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
