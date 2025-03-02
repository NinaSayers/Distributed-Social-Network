package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
)

type Post struct {
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Create(post *dto.CreatePostDTO) (*Post, error) {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM user WHERE user_id = ?", post.UserID).Scan(&count)
	if err != nil {
		return nil, NewErrUserCheck(err)
	}
	if count == 0 {
		return nil, ErrNoRecord
	}

	stmt := `INSERT INTO post (user_id, post_id, content) VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, post.UserID, post.PostID, post.Content)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return m.Get(post.PostID)
}

func (m *PostModel) Get(id string) (*Post, error) {
	stmt := `SELECT post_id, user_id, content, created_at, updated_at FROM post WHERE post_id = ?`
	row := m.DB.QueryRow(stmt, id)

	p := &Post{}
	err := row.Scan(&p.PostID, &p.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return p, nil
}

func (m *PostModel) HGet(id string) (*dto.HPost, error) {
	stmt := `
        SELECT 
            p.post_id, p.user_id, p.content, p.created_at, p.updated_at,
            u.username, u.email, u.name, u.avatar, u.created_at
        FROM 
            post p
        JOIN 
            user u ON p.user_id = u.user_id
        WHERE 
            p.post_id = ?`
	row := m.DB.QueryRow(stmt, id)

	p := &dto.HPost{}
	u := &dto.CoreUserDTO{}
	err := row.Scan(
		&p.PostID, &u.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt, &u.UserName,
		&u.Email, &u.Name, &u.Avatar, &u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, NewErrDatabaseOperationFailed(err)
	}
	p.User = *u

	return p, nil
}

func (m *PostModel) ListByUser(userID string) ([]*dto.HPost, error) {

	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM user WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}
	if count == 0 {
		return nil, ErrNoRecord
	}
	stmt := `
        SELECT 
            p.post_id, p.user_id, p.content, p.created_at, p.updated_at,
            u.username, u.email, u.name, u.avatar, u.created_at
        FROM 
            post p
        JOIN 
            user u ON p.user_id = u.user_id
        WHERE 
            p.user_id = ?
		ORDER BY p.created_at DESC`

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	defer rows.Close()

	posts := []*dto.HPost{}

	for rows.Next() {
		p := &dto.HPost{}
		u := &dto.CoreUserDTO{}
		err := rows.Scan(
			&p.PostID, &u.UserID, &p.Content, &p.CreatedAt, &p.UpdatedAt, &u.UserName,
			&u.Email, &u.Name, &u.Avatar, &u.CreatedAt,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrNoRecord
			}
			return nil, NewErrDatabaseOperationFailed(err)
		}
		p.User = *u
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, NewErrDatabaseOperationFailed(err)
	}

	return posts, nil
}

func (m *PostModel) Delete(postID string) error {

	res, err := m.DB.Exec("DELETE FROM post WHERE post_id = ?", postID)
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
