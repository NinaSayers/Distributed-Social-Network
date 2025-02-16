package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
)

type Message struct {
	MessageID string    `json:"message_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessageModel struct {
	DB *sql.DB
}

func (m *MessageModel) Create(message *dto.CreateMessageDTO) (int, error) {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE user_id = ?", message.UserID).Scan(&count)
	if err != nil {
		return 0, NewErrUserCheck(err)
	}
	if count == 0 {
		return 0, ErrNoRecord
	}

	stmt := `INSERT INTO messages (user_id, message_id, content) VALUES (?, ?, ?)`
	result, err := m.DB.Exec(stmt, message.UserID, message.MessageID, message.Content)
	if err != nil {
		return 0, NewErrDatabaseOperationFailed(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, NewErrDatabaseOperationFailed(err)
	}

	return int(id), nil
}

func (m *MessageModel) Get(id string) (*Message, error) {
	stmt := `SELECT message_id, user_id, content, created_at, updated_at FROM messages WHERE message_id = ?`
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

func (m *MessageModel) ListByUser(userID string) ([]*Message, error) {

	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}
	if count == 0 {
		return nil, ErrNoRecord
	}

	stmt := `SELECT message_id, user_id, content, created_at 
	 		FROM messages
	 		WHERE user_id = ?
	 		ORDER BY created_at DESC`

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		msg := &Message{}
		err = rows.Scan(&msg.MessageID, &msg.UserID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, NewErrDatabaseOperationFailed(err)
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return messages, nil
}

func (m *MessageModel) Delete(messageID string) error {

	res, err := m.DB.Exec("DELETE FROM messages WHERE message_id = ?", messageID)
	if err != nil {
		return NewErrDatabaseOperationFailed(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return NewErrDatabaseOperationFailed(err)
	}

	if rowsAffected == 0 {
		return ErrNoRecord
	}

	return nil
}
