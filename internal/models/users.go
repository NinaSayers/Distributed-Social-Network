package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	user_id       int
	username      string
	email         string
	password_hash string
	created_at    time.Time
	updated_at    time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(username, email, password string) (int, error) {
	stmt := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec(stmt, username, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *UserModel) Get(id int) (*User, error) {
	return nil, nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {

	var id int
	var hashedPassword []byte

	stmt := "SELECT id, hashed_password FROM users WHERE email = ?"
	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, ErrInvalidCredentials
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return int(id), nil
}
