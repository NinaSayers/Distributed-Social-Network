package dto

import "time"

type AuthUserDTO struct {
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"username"`
	Email    string `json:"email"`

	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
