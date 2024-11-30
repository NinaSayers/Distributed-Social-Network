package models

import (
	"database/sql"
	"time"
	"fmt"
)

type Post struct {
	PostID	   int		`json:"post_id"`
	UserID     int		`json:"user_id"`
	Content    string	`json:"content"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAT time.Time	`json:"updated_at"`
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Create(user_id int, content string) (int, error){

	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", user_id).Scan(&count)
	if err != nil {
	 return 0, fmt.Errorf("error checking user existence: %w", err)
	}
	if count == 0 {
	 return 0, fmt.Errorf("user with ID %d does not exist", user_id)
	}
	
	stmt := `INSERT INTO posts (user_id, content) VALUES (?, ?)`
	
	result, err := m.DB.Exec(stmt, user_id, content)
	if err != nil {
	 return 0, fmt.Errorf("error creating post: %w", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
	 return 0, fmt.Errorf("error getting post ID: %w", err)
	}
	
	return int(id), nil
}


func (m *PostModel) Get(id int) (*Post, error){
	return nil, nil
}