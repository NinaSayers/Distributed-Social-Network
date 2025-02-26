package models

import "database/sql"

func CheckRepostExistence(userID string, messageID string, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM repost WHERE user_id = ? AND post_id = ?", userID, messageID).Scan(&count)
	if err != nil {
		return 0, NewErrDatabaseOperationFailed(err)
	}
	return count, nil
}
