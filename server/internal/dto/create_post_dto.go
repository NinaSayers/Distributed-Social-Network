package dto

type CreatePostDTO struct {
	PostID  string `json:"message_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
