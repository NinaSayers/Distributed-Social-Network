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

func (app *Application) showProfile(userID string) (*User, error) {
    user, err := app.service.GetUser(userID)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (app *Application) showMyProfile() (*User, error) {
    user, err := app.service.GetUser(app.user.UserID)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (app *Application) updateProfile(username, email string) error {
    err := app.service.UpdateUser(app.user.UserID, username, email)
    if err != nil {
        return err
    }
    return nil
}
