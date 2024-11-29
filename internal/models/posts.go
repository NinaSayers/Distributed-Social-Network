package models

import (
	"database/sql"
	"time"
)

type Post struct {
	post_id	   int
	user_id    int
	content    string
	created_at time.Time
	updated_at time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Create(user_id int, content string) (int, error){
	return 0, nil
}

func (m *PostModel) Get(id int) (*Post, error){
	return nil, nil
}