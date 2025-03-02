package dto

import "time"

type CreatePostDTO struct {
	PostID    string    `json:"post_id,omitempty"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
