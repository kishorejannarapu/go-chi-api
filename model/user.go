package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"-"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
