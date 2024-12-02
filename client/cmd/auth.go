package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	UserID    int       `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	User  `json:"user"`
	Token string `json:"token"`
}

func signUp() {
	var username, email, password string
	fmt.Print("Nombre de usuario: ")
	fmt.Scan(&username)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	user := map[string]string{"username": username, "email": email, "password": password}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuario creado con éxito")
}

func login() (*Client, error) {
	var email, password string
	fmt.Print("Correo de usuario: ")
	fmt.Scan(&email)
	fmt.Print("Contraseña: ")
	fmt.Scan(&password)

	credentials := map[string]string{"Email": email, "password": password}
	jsonData, err := json.Marshal(credentials)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var response Client
	err = json.Unmarshal(body, &response)

	fmt.Println("Bienvenido %s", response.UserName)
	return &response, nil
}
