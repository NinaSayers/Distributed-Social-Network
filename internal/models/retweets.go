package models

import (
	"database/sql"
	"time"
)

type Retweet struct {
	RetweetID int       `json:"retweet_id"`
	UserID    int       `json:"user_id"`
	MessageID int       `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RetweetModel struct {
	DB *sql.DB
}

// RetweetModel methods
func (m *RetweetModel) CreateRetweet(userID, messageID int) error {
	err := CheckUserExistence(userID, m.DB)
	if err != nil {
	 return err
	}
   
	err = CheckMessageExistence(messageID, m.DB)
	if err != nil {
	 return err
	}
   
	exist, err := CheckRetweetExistence(userID, messageID, m.DB)
	if err != nil {
	 	return err
	}
	if exist > 0 {
	 	return ErrRelationshipExists
	}
   
	stmt := `INSERT INTO retweets (user_id, message_id, created_at) VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, userID, messageID, time.Now())
	if err != nil {
	 	return NewErrDatabaseOperationFailed(err)
	}
   
	return nil
}

func (m *RetweetModel) UndoRetweet(userID, messageID int) error {

	exists, err := CheckRetweetExistence(userID, messageID, m.DB)
	if err != nil {
	 	return err
	}
	if exists == 0 {
	 	return ErrNoRecord
	}

	stmt := `DELETE FROM retweets WHERE user_id = ? AND message_id = ?`
	result, err := m.DB.Exec(stmt, userID, messageID)
	if err != nil {
	 	return NewErrDatabaseOperationFailed(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
	 	return NewErrDatabaseOperationFailed(err)
	}
	if rowsAffected == 0 {
	 	return ErrNoRecord
	}
   
	return nil
}