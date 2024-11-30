package models

import "database/sql"

func CheckUserExistence(userID int, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userID).Scan(&count)
	if err != nil {
	 	return NewErrUserCheck(err)
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}