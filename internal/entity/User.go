package entity

import "time"

type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	HashPassword string    `json:"password"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
}
