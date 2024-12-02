package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Tweet struct {
	Content   string    `json:"content"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	TweetId   string    `json:"message_id"`
}

func listUsers() {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var followers []User
	err = json.Unmarshal(body, &followers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, follower := range followers {
		fmt.Printf("ID: %s, Username: %s, Email: %s\n", follower.UserID, follower.UserName, follower.Email)
	}
}

func getUser() {
	var id string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	resp, err := http.Get(baseURL + "/users/" + id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var response User
	err = json.Unmarshal(body, &response)
	displayProfileHeader(response)
}

func displayProfileHeader(user User) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("  %s (%s)\n", user.UserName, user.Email)
	fmt.Printf("  🗓️  Joined %s\n", user.CreatedAt.Local())
	// fmt.Printf("  📊 %d Following   👥 %d Followers\n", user.Following, user.Followers)
	fmt.Println(strings.Repeat("=", 50))
}

func listFollowers() {
	var id string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	resp, err := http.Get(baseURL + "/users/" + id + "/followers")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var followers []User
	err = json.Unmarshal(body, &followers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, follower := range followers {
		fmt.Printf("ID: %s, Username: %s, Email: %s\n", follower.UserID, follower.UserName, follower.Email)
	}
}

func listFollowing() {
	var id string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	resp, err := http.Get(baseURL + "/users/" + id + "/following")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var followings []User
	err = json.Unmarshal(body, &followings)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, followings := range followings {
		fmt.Printf("ID: %s, Username: %s, Email: %s\n", followings.UserID, followings.UserName, followings.Email)
	}
}
