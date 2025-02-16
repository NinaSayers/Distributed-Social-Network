package dto

type CreateMessageDTO struct {
	MessageID string `json:"message_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
}
