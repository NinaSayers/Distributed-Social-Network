package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
func (s *Service) CreateUser(username, email, password string) (string, error) {
	payload := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(payload)

	resp, err := s.client.Post(baseURL+"/users", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create user: status %d", resp.StatusCode)
	}

	var response struct {
		UserID string `json:"user_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	return response.UserID, nil
}

// GetUser retrieves a user by ID
func (s *Service) GetUser(userID string) (*User, error) {
	resp, err := s.client.Get(baseURL + "/users/" + userID)
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
func (s *Service) DeleteUser(userID string) error {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+userID, nil)
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

func (s *Service) UpdateUser(userID, username, email string) error {
	payload := map[string]string{
		"username": username,
		"email":    email,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPut, baseURL+"/users/"+userID, bytes.NewBuffer(body))
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

	body, err := json.Marshal(payload)
	// jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	resp, err := http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response Client

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &response, nil
}

// func (s *Service) Logout() error {
//     req, err := http.NewRequest(http.MethodPost, baseURL+"/auth/logout", nil)
//     if err != nil {
//         return fmt.Errorf("error creating logout request: %w", err)
//     }

//     // Si estás usando un token JWT, agrégalo al header de la solicitud.
//     if s.token != "" {
//         req.Header.Set("Authorization", "Bearer "+s.token)
//     }

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         return fmt.Errorf("error sending logout request: %w", err)
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         return fmt.Errorf("logout failed with status code: %d", resp.StatusCode)
//     }

//     return nil
// }
////////////////////

// CreateMessage sends a message for a user
func (s *Service) CreateMessage(userID string, content string) (*Message, error) {
	var payload = struct {
		UserID  string `json:"user_id"`
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
func (s *Service) ListUserMessages(userID string) ([]Message, error) {
	resp, err := s.client.Get(baseURL + "/users/" + userID + "/messages")
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
func (s *Service) FollowUser(followerID, followeeID string) error {
	payload := map[string]string{
		"follower_id": followerID,
		"followee_id": followeeID,
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/users/"+followeeID+"/follow", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to follow user: status %d", resp.StatusCode)
	}
	return nil
}

func (s *Service) UnfollowUser(userID, followeeID string) error {
	payload := map[string]string{
		"user_id": userID,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+followeeID+"/follow", bytes.NewBuffer(body))
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

func (s *Service) ListFollowers(user_id string) ([]User, error) {

	resp, err := http.Get(baseURL + "/users/" + user_id + "/followers")
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

func (s *Service) ListFollowing(user_id string) ([]User, error) {

	resp, err := http.Get(baseURL + "/users/" + user_id + "/following")
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
func (s *Service) GetMessage(messageID string) (*Message, error) {
    // Realizar una solicitud HTTP GET al servidor para obtener el mensaje
    resp, err := s.client.Get(baseURL + "/messages/" + messageID)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to retrieve message: status %d", resp.StatusCode)
    }

    // Decodificar la respuesta JSON en un mensaje
    var message Message
    if err := json.NewDecoder(resp.Body).Decode(&message); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return &message, nil
}

func (s *Service) GetUserMessages(userID string) ([]Message, error) {
    // Realizar una solicitud HTTP GET al servidor para obtener los mensajes del usuario
    resp, err := s.client.Get(baseURL + "/users/" + userID + "/messages")
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to retrieve messages: status %d", resp.StatusCode)
    }

    // Decodificar la respuesta JSON en una lista de mensajes
    var messages []Message
    if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {
        return nil, fmt.Errorf("failed to decode response: %w", err)
    }

    return messages, nil
}
