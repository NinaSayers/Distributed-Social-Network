package dto

import "time"

type HPost struct {
	PostID    string    `json:"post_id"`
	User      U         `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
