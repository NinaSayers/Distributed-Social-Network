package models

import "database/sql"

func CheckRetweetExistence(userID int, messageID int, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM retweets WHERE user_id = ? AND message_id = ?", userID, messageID).Scan(&count)
	if err != nil {
		return 0, NewErrDatabaseOperationFailed(err)
	}
	return count, nil
}