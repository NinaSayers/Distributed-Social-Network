package dto

import "time"

type AuthUserDTO struct {
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"username"`
	Email    string `json:"email"`

	PasswordHash string `json:"password_hash"`

	Bio    string `json:"bio"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Following int `json:"following,omitempty"`
	Followers int `json:"followers,omitempty"`
}
