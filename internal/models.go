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

type Message struct {
	message_id int
	user_id    int
	content    string
	created_at time.Time
	updated_at time.Time
}
