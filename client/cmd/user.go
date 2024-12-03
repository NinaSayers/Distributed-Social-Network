package main

import (
	"fmt"
)

func (app *Application) listUsers() {

	users, err := app.service.ListUsers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	displayUsers(users)
}

func (app *Application) showProfile() {
	var id int
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	user, err := app.service.GetUser(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	messages, err := app.service.ListUserMessages(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayProfile(*user, messages)
}

func (app *Application) showMyProfile() {

	user, err := app.service.GetUser(app.user.UserID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	messages, err := app.service.ListUserMessages(app.user.UserID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayProfile(*user, messages)
}

func (app *Application) updateProfile() {
	var username, email string
	fmt.Print("Nuevo nombre de usuario: ")
	fmt.Scan(&username)
	fmt.Print("Nuevo email: ")
	fmt.Scan(&email)

	err := app.service.UpdateUser(app.user.UserID, username, email)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Perfil actualizado con éxito")
}
