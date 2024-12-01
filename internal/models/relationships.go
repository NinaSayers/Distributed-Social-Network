package models

import (
	"database/sql"
	"time"
	"errors"
)

type Relationship struct {
	RelationshipID	int 		`json:"relationship_id"`
	FollowerID 		int 		`json:"follower_id"`
	FolloweeID 		int 		`json:"followee_id"`
	CreatedAt 		time.Time 	`json:"created_at"`
}

type RelationshipModel struct {
	DB *sql.DB
}

func (m *RelationshipModel) FollowUser(followerID, followeeID int) error {

	err := CheckUserExistence(followerID, m.DB)
	if err != nil {
	 	return err
	}
   
	err = CheckUserExistence(followeeID, m.DB)
	if err != nil {
	 	return err
	}	
   
	err = CheckRelationshipExistence(followerID, followeeID, m.DB)
	if err != nil {
		return err
    }
   
	stmt := `INSERT INTO relationships (follower_id, followee_id, created_at) VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, followerID, followeeID, time.Now())
	if err != nil {
	 	return NewErrDatabaseOperationFailed(err)
	}
   
	return nil
}

func (m *RelationshipModel) UnfollowUser(relationshipID int) error {

	err := CheckRelationshipExistenceByID(relationshipID, m.DB)
	if err != nil {
	 	return err
	}
   
	stmt := `DELETE FROM relationships WHERE relationship_id = ?`
	result, err := m.DB.Exec(stmt, relationshipID)
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

func (m *RelationshipModel) ListFollowers(userID int) ([]*Relationship, error) {

	err := CheckUserExistenceAsFollowee(userID, m.DB)

	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			return []*Relationship{}, nil
		}
		return nil, err
	}

 	stmt := `SELECT follower_id, FROM relationships WHERE followee_id = ?`
 	rows, err := m.DB.Query(stmt, userID)
 	if err != nil {
 	 	return nil, NewErrDatabaseOperationFailed(err)
 	}
 	defer rows.Close()

 	relationships := []*Relationship{}
 	for rows.Next() {
 	 	r := &Relationship{}
 	 	err := rows.Scan(&r.RelationshipID, &r.FollowerID, &r.FolloweeID, &r.CreatedAt)
 	 	if err != nil {
 	  		return nil, NewErrDatabaseOperationFailed(err)
 	 	}
 	 	relationships = append(relationships, r)
 	}

 	if err = rows.Err(); err != nil {
 	 	return nil, NewErrDatabaseOperationFailed(err)
 	}

 	return relationships, nil
}

func (m *RelationshipModel) ListFollowing(userID int) ([]*Relationship, error) {

	err := CheckUserExistenceAsFollower(userID, m.DB)

	if err != nil {
		if errors.Is(err, ErrNoRecord) {
			return []*Relationship{}, nil
		}
		return nil, err
	}

 	stmt := `SELECT followee_id, FROM relationships WHERE follower_id = ?`
 	rows, err := m.DB.Query(stmt, userID)
 	if err != nil {
 	 	return nil, NewErrDatabaseOperationFailed(err)
 	}
 	defer rows.Close()

 	relationships := []*Relationship{}
 	for rows.Next() {
 	 	r := &Relationship{}
 	 	err := rows.Scan(&r.RelationshipID, &r.FollowerID, &r.FolloweeID, &r.CreatedAt)
 	 	if err != nil {
 	  		return nil, NewErrDatabaseOperationFailed(err)
 	 	}
 	 	relationships = append(relationships, r)
 	}

 	if err = rows.Err(); err != nil {
 	 	return nil, NewErrDatabaseOperationFailed(err)
 	}

 	return relationships, nil
}