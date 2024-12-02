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

func displayPosts(tweets []Tweet) {
	fmt.Println("Posts:")
	for i, tweet := range tweets {
		if i == 5 {
			break
		}
		fmt.Printf("\n%s: %s\n", tweet.UserId, tweet.Content)
		fmt.Printf("  📅 %s\n", tweet.CreatedAt)
	}
	fmt.Println(strings.Repeat("-", 50))
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

	fmt.Println(string(body))
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

	fmt.Println(string(body))
}
