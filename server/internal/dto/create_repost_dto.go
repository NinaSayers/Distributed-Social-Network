package dto

type CreateRepostDTO struct {
	RepostID string `json:"repost_id,omitempty"`
	PostID   string `json:"post_id"`
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
}
