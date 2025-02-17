package dto

type CreatePostDTO struct {
	PostID  string `json:"message_id,omitempty"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
