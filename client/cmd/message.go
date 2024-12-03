package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (app *Application) createMessageComponent() {
	var content string
	fmt.Print("Contenido del mensaje: ")
	fmt.Scan(&content)

	message, err := app.service.CreateMessage(app.user.UserID, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayPost(*message, *app.user)
}

func (app *Application) getMessage() {
	var id string
	fmt.Print("ID del mensaje: ")
	fmt.Scan(&id)

	message, err := app.service.GetMessage(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayPosts([]Message{*message})
}

func deleteMessage() {
	var id string
	fmt.Print("ID del mensaje: ")
	fmt.Scan(&id)

	req, err := http.NewRequest(http.MethodDelete, baseURL+"/messages/"+id, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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
