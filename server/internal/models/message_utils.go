package models

import "database/sql"

func CheckMessageExistence(messageID string, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM post WHERE post_id = ?", messageID).Scan(&count)
	if err != nil {
		return ErrMessageCheck
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}
