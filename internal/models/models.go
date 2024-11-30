package models

import "database/sql"

type Models struct {
	User *UserModel
	Post *PostModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User: &UserModel{DB: db},
		Post: &PostModel{DB: db},
	}
}
