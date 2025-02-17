package models

import "database/sql"

func CheckRelationshipExistence(followerID string, followeeID string, db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM follow WHERE user_id = ? AND followee_id = ?", followerID, followeeID).Scan(&count)
	if err != nil {
		return 0, NewErrDatabaseOperationFailed(err)
	}
	return count, nil
}

func CheckRelationshipExistenceByID(relationshipID int, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM relationships WHERE id = ?", relationshipID).Scan(&count)
	if err != nil {
		return NewErrUserCheck(err)
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}

func CheckUserExistenceAsFollower(userID string, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM follow WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return NewErrUserCheck(err)
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}

func CheckUserExistenceAsFollowee(userID int, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM relationships WHERE followee_id = ?", userID).Scan(&count)
	if err != nil {
		return NewErrUserCheck(err)
	}
	if count == 0 {
		return ErrNoRecord
	}
	return nil
}
