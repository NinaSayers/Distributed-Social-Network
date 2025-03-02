package dto

import "time"

type CoreUserDTO struct {
	UserID    string    `json:"user_id,omitempty"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
