package entity

import "time"

type Project struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Owner       int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}
