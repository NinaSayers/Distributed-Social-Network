package models

import "database/sql"

func CheckMessageExistence(messageID int, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM messages WHERE message_id = ?", messageID).Scan(&count)
	if err != nil {
		return ErrMessageCheck
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}
