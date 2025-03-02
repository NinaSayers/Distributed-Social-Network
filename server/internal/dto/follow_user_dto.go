package dto

import "time"

type FollowUserDTO struct {
	FollowId   string    `json:"follow_id,omitempty"`
	UserID     string    `json:"user_id"`
	FolloweeID string    `json:"followee_id"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
