package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Service struct to handle API requests
type Service struct {
	client *http.Client
}

// NewService creates a new service instance
func NewService() *Service {
	return &Service{client: &http.Client{}}
}

// CreateUser sends a request to create a new user
func (s *Service) CreateUser(username, email, password string) (int, error) {
	payload := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(payload)

	resp, err := s.client.Post(baseURL+"/users", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("failed to create user: status %d", resp.StatusCode)
	}

	var response struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}
	return response.UserID, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(userID int) (*User, error) {
	resp, err := s.client.Get(baseURL + "/users/" + strconv.Itoa(userID))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve user: status %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &user, nil
}

// ListUsers retrieves all users
func (s *Service) ListUsers() ([]User, error) {
	resp, err := s.client.Get(baseURL + "/users")
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve users: status %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return users, nil
}

// DeleteUser deletes a user by ID
func (s *Service) DeleteUser(userID int) error {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+strconv.Itoa(userID), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete user: status %d", resp.StatusCode)
	}
	return nil
}

func (s *Service) UpdateUser(userID int, username, email string) error {
	payload := map[string]string{
		"username": username,
		"email":    email,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPut, baseURL+"/users/"+strconv.Itoa(userID), bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user: status %d", resp.StatusCode)
	}
	return nil
}

// Login sends a request to authenticate a user
func (s *Service) Login(email, password string) (*Client, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}

	body, _ := json.Marshal(payload)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	resp, err := http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response Client
	err = json.Unmarshal(body, &response)

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &response, nil
}

////////////////////

// CreateMessage sends a message for a user
func (s *Service) CreateMessage(userID int, content string) (*Message, error) {
	var payload = struct {
		UserID  int    `json:"user_id"`
		Content string `json:"content"`
	}{userID, content}

	body, _ := json.Marshal(payload)

	resp, err := s.client.Post(baseURL+"/messages", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create message: status %d", resp.StatusCode)
	}

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &message, nil
}

// ListUserMessages retrieves messages for a user
func (s *Service) ListUserMessages(userID int) ([]Message, error) {
	resp, err := s.client.Get(baseURL + "/users/" + strconv.Itoa(userID) + "/messages")
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve messages: status %d", resp.StatusCode)
	}

	var messages []Message
	if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return messages, nil
}

// FollowUser sends a request to follow a user
func (s *Service) FollowUser(followerID, followeeID int) error {
	payload := map[string]int{
		"follower_id": followerID,
		"followee_id": followeeID,
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/users/"+strconv.Itoa(followeeID)+"/follow", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to follow user: status %d", resp.StatusCode)
	}
	return nil
}

func (s *Service) UnfollowUser(userID, followeeID int) error {
	payload := map[string]int{
		"user_id": userID,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+strconv.Itoa(followeeID)+"/follow", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to unfollow user: status %d", resp.StatusCode)
	}
	return nil
}

func (s *Service) ListFollowers(user_id int) ([]User, error) {
	
	resp, err := http.Get(baseURL + "/users/" + strconv.Itoa(user_id) + "/followers")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve users: status %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return users, nil

}

func (s *Service) ListFollowing(user_id int) ([]User, error) {
	
	resp, err := http.Get(baseURL + "/users/" + strconv.Itoa(user_id) + "/following")
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve users: status %d", resp.StatusCode)
	}

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return users, nil

}
func (s *Service) GetMessage(id int) (*Message, error) {
	id_string := strconv.Itoa(id)
	resp, err := s.client.Get(baseURL + "/messages/" + id_string)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve message: status %d", resp.StatusCode)
	}

	var message Message
	if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &message, nil
}
