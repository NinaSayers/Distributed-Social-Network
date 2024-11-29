package models

import (
	"database/sql"
	"time"
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

func (m *UserModel) Create(username, email, password_hash string) (int, error) {
	stmt := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
   
	result, err := m.DB.Exec(stmt, username, email, password_hash)
	if err != nil {
	 return 0, err
	}
   
	id, err := result.LastInsertId()
	if err != nil {
	 return 0, err
	}
   
	return int(id), nil
   }

func (m *UserModel) Get(id int) (*User, error){
	return nil, nil
}


