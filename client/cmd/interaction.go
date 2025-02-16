package main

import (
	"fmt"
)

func (app *Application) followUser() {
	var id int
	fmt.Print("ID de usuario a seguir: ")
	fmt.Scan(&id)

	err := app.service.FollowUser(app.user.UserID, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Ahora sigues a este usuario")
}

func (app *Application) unfollowUser() {
	var id int
	fmt.Print("ID de usuario a dejar de seguir: ")
	fmt.Scan(&id)

	err := app.service.UnfollowUser(app.user.UserID, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// paylaod.UserId = client.UserID

	fmt.Println("Ahora no sigues a este usuario")
}

func (app *Application) listFollowers() {
	var id int
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	users, err := app.service.ListFollowers(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	displayUsers(users)

}

func (app *Application) listFollowing() {
	var id int
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	users, err := app.service.ListFollowing(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	displayUsers(users)

}
