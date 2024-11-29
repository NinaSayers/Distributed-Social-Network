package main

import (
	"net/http"
	"fmt"
)

func (app *application) feed(w http.ResponseWriter, r *http.Request) { // creo que esto hay que borrarlo
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Hello from Distnet"))
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a post"))
}

func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	username := "Mike"
	email := "mike@example.com"
	password_hash := "abcd1234"

	id, err := app.users.Create(username, email, password_hash)
	
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/view?id=%d", id), http.StatusSeeOther) // verify route!!!!!!!!!!!!!!!!

	// w.Write([]byte("Creating user"))
}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

func (app *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

func (app *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListFollowersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListFollowingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) CreateTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) GetTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListUserTweetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) GetTimelineHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) RetweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) UndoRetweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) FavoriteTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) UnfavoriteTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListReceivedMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListSentMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
