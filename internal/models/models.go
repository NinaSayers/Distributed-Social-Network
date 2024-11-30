package models

import "database/sql"

type Models struct {
	User 		 *UserModel
	Message 	 *MessageModel
	Relationship *RelationshipModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User: 		  &UserModel{DB: db},
		Message: 	  &MessageModel{DB: db},
		Relationship: &RelationshipModel{DB: db},
	}
}
