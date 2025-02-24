package main

// import (
// 	"fmt"
// )

func (app *Application) followUser(userID string) error {
    err := app.service.FollowUser(app.user.UserID, userID)
    if err != nil {
        return err
    }
    return nil
}

func (app *Application) unfollowUser(userID string) error {
    err := app.service.UnfollowUser(app.user.UserID, userID)
    if err != nil {
        return err
    }
    return nil
}

func (app *Application) listFollowers(userID string) ([]User, error) {
    users, err := app.service.ListFollowers(userID)
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (app *Application) listFollowing(userID string) ([]User, error) {
    users, err := app.service.ListFollowing(userID)
    if err != nil {
        return nil, err
    }
    return users, nil
}
