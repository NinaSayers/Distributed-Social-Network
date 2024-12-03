package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (app *Application) followUser() {
	var id string
	fmt.Print("ID de usuario a seguir: ")
	fmt.Scan(&id)

	var payload struct {
		FollowerID int `json:"follower_id"`
		FolloweeID int `json:"followee_id"`
	}

	// payload.FollowerID = client.UserID
	followeeID, err := strconv.Atoi(id)
	payload.FolloweeID = followeeID

	if err != nil {
		fmt.Println("Error: ID de usuario inválido")
		return
	}
	jsonData, err := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/users/"+id+"/follow", "application/json", bytes.NewBuffer(jsonData))
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

func (app *Application) unfollowUser() {
	var id string
	fmt.Print("ID de usuario a dejar de seguir: ")
	fmt.Scan(&id)
	var paylaod struct {
		UserId int `json:"user_id"`
	}
	// paylaod.UserId = client.UserID
	jsonData, err := json.Marshal(paylaod)
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+id+"/follow", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
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
