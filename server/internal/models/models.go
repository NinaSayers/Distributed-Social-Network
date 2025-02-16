package models

import "database/sql"

type Models struct {
	User         *UserModel
	Post         *PostModel
	Relationship *RelationshipModel
	Retweet      *RetweetModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User:         &UserModel{DB: db},
		Post:         &PostModel{DB: db},
		Relationship: &RelationshipModel{DB: db},
		Retweet:      &RetweetModel{DB: db},
	}
}
