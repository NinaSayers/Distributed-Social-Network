package models

import "database/sql"

func CheckUserExistence(userID string, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return NewErrUserCheck(err)
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}
