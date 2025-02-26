package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
)

type Repost struct {
	RepostID  string    `json:"repost_id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RepostModel struct {
	DB *sql.DB
}

// RepostModel methods
func (m *RepostModel) CreateRepost(repost *dto.CreateRepostDTO) (*Repost, error) {
	err := CheckUserExistence(repost.UserID, m.DB)
	if err != nil {
		return nil, err
	}

	// err = CheckMessageExistence(repost.PostID, m.DB)
	// if err != nil {
	// 	return err
	// }

	exist, err := CheckRepostExistence(repost.UserID, repost.PostID, m.DB)
	if err != nil {
		return nil, err
	}
	if exist > 0 {
		return nil, ErrRelationshipExists
	}

	stmt := `INSERT INTO repost (repost_id, user_id, post_id, created_at) VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, repost.RepostID, repost.UserID, repost.PostID, time.Now())
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return m.Get(repost.RepostID)
}

func (m *RepostModel) UndoRepost(userID, postID string) error {

	exists, err := CheckRepostExistence(userID, postID, m.DB)
	if err != nil {
		return err
	}
	if exists == 0 {
		return ErrNoRecord
	}

	stmt := `DELETE FROM repost WHERE user_id = ? AND post_id = ?`
	result, err := m.DB.Exec(stmt, userID, postID)
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

func (m *RepostModel) Get(id string) (*Repost, error) {
	stmt := `SELECT repost_id, post_id, user_id, created_at, updated_at FROM post WHERE repost_id = ?`
	row := m.DB.QueryRow(stmt, id)

	p := &Repost{}
	err := row.Scan(&p.RepostID, &p.PostID, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return p, nil
}

func (m *RepostModel) Delete(repostID string) error {

	res, err := m.DB.Exec("DELETE FROM post WHERE repost_id = ?", repostID)
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
