package dto

type FollowUserDTO struct {
	FollowId   string `json:"follow_id,omitempty"`
	UserID     string `json:"user_id"`
	FolloweeID string `json:"followee_id"`
}
