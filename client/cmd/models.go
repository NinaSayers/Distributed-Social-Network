package main

import "time"

type Message struct {
	MessageID int       `json:"message_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Relationship struct {
	RelationshipID int       `json:"relationship_id"`
	FollowerID     int       `json:"follower_id"`
	FolloweeID     int       `json:"followee_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Retweet struct {
	RetweetID int       `json:"retweet_id"`
	UserID    int       `json:"user_id"`
	MessageID int       `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
	Message   Message   `json:"message"`
}

type User struct {
	UserID    int       `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Client struct {
	User  `json:"user"`
	Token string `json:"token"`
}
