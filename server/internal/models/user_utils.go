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

func CheckFollowers(userID string, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM follow WHERE followee_id = ?", userID).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func CheckFollowing(userID string, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM follow WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
