package dto

type CreateUserDTO struct {
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
