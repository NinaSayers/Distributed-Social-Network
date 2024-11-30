package models

import (
	"database/sql"
	"time"
	"errors"
)

type Message struct {
	MessageID  int		`json:"message_id"`
	UserID     int		`json:"user_id"`
	Content    string	`json:"content"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type MessageModel struct {
	DB *sql.DB
}

func (m *MessageModel) Create(user_id int, content string) (int, error) {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", user_id).Scan(&count)
	if err != nil {
	 	return 0, NewErrUserCheck(err)
	}
	if count == 0 {
	 	return 0, ErrNoRecord
	}
   
	stmt := `INSERT INTO messages (user_id, content) VALUES (?, ?)`
	result, err := m.DB.Exec(stmt, user_id, content)
	if err != nil {
	 	return 0, NewErrDatabaseOperationFailed(err) 
	}
   
	id, err := result.LastInsertId()
	if err != nil {
	 	return 0, NewErrDatabaseOperationFailed(err) 
	}
   
	return int(id), nil
}
   
func (m *MessageModel) Get(id int) (*Message, error) {
	stmt := `SELECT id, user_id, content, created_at, updated_at FROM messages WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
   
	p := &Message{}
	err := row.Scan(&p.MessageID, &p.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
	 	if errors.Is(err, sql.ErrNoRows) {
	  		return nil, ErrNoRecord
	 	}
	 	return nil, NewErrDatabaseOperationFailed(err)
	}
   
	return p, nil
}