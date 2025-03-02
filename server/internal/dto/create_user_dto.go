package dto

type CreateUserDTO struct {
	UserID       string `json:"user_id,omitempty"`
	UserName     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash,omitempty"`
	Bio          string `json:"bio,omitempty"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar,omitempty"`
}
