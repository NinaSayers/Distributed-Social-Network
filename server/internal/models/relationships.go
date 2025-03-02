package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
)

type Follow struct {
	FollowID   string    `json:"follow_id"`
	UserID     string    `json:"user_id"`
	FolloweeID string    `json:"followee_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type RelationshipModel struct {
	DB *sql.DB
}

func (m *RelationshipModel) FollowUser(follow *dto.FollowUserDTO) error {

	err := CheckUserExistence(follow.UserID, m.DB)
	if err != nil {
		return err
	}

	// err = CheckUserExistence(followeeID, m.DB)
	// if err != nil {
	// 	return err
	// }

	exist, err := CheckRelationshipExistence(follow.UserID, follow.FolloweeID, m.DB)
	if err != nil {
		return err
	}
	if exist > 0 {
		return ErrRelationshipExists
	}

	stmt := `INSERT INTO follow (follow_id, user_id, followee_id) VALUES (?, ?, ?)`
	if follow.CreatedAt.IsZero() {
		follow.CreatedAt = time.Now()
	}

	if follow.UpdatedAt.IsZero() {
		follow.UpdatedAt = time.Now()
	}

	_, err = m.DB.Exec(stmt, follow.FollowId, follow.UserID, follow.FolloweeID, time.Now())
	if err != nil {
		return NewErrDatabaseOperationFailed(err)
	}

	return nil
}

func (m *RelationshipModel) Get(id string) (*Follow, error) {
	stmt := `SELECT follow_id, user_id, followee_id, created_at FROM follow WHERE follow_id = ?`
	row := m.DB.QueryRow(stmt, id)

	f := &Follow{}
	err := row.Scan(&f.FollowID, &f.UserID, &f.FolloweeID, &f.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return f, nil
}

func (m *RelationshipModel) UnfollowUser(userId, followeeId string) error {
	exists, err := CheckRelationshipExistence(userId, followeeId, m.DB)
	if err != nil {
		return err
	}
	if exists == 0 {
		return ErrNoRecord
	}

	stmt := `DELETE FROM relationships WHERE follower_id = ? AND followee_id = ?`
	result, err := m.DB.Exec(stmt, userId, followeeId)
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

func (m *RelationshipModel) ListFollowing(userID string) ([]string, error) {

	err := CheckUserExistenceAsFollower(userID, m.DB)

	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			return []string{}, nil
		}
		return nil, err
	}

	stmt := `
		SELECT followee_id
		FROM follow
		WHERE user_id = ?
		`
	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}
	defer rows.Close()

	users := []string{}
	for rows.Next() {
		var u string
		err := rows.Scan(&u)
		if err != nil {
			return nil, NewErrDatabaseOperationFailed(err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return users, nil
}
func (m *RelationshipModel) ListFollowers(userID string) ([]string, error) {

	err := CheckUserExistenceAsFollowee(userID, m.DB)

	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			return []string{}, nil
		}
		return nil, err
	}

	stmt := `
		SELECT user_id
		FROM follow
		WHERE followee_id = ?
		`
	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}
	defer rows.Close()

	users := []string{}
	for rows.Next() {
		var u string
		err := rows.Scan(&u)
		if err != nil {
			return nil, NewErrDatabaseOperationFailed(err)
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return users, nil
}

func (m *RelationshipModel) Delete(followId string) error {

	res, err := m.DB.Exec("DELETE FROM follow WHERE follow_id = ?", followId)
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
