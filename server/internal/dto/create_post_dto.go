package dto

type CreatePostDTO struct {
	PostID  string `json:"post_id,omitempty"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
