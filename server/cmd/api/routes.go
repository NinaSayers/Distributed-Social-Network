package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", app.feed)
	// mux.HandleFunc("GET	/users", app.ListUsersHandler)
	mux.HandleFunc("POST /users", app.CreateUserHandler)
	mux.HandleFunc("GET /users/{id}", app.GetUserHandler)
	mux.HandleFunc("GET /users/email/{mail}", app.GetUserEmailHandler)
	// mux.HandleFunc("PUT /users/{id}", app.UpdateUserHandler)
	// mux.HandleFunc("DELETE /users/{id}", app.DeleteUserHandler)

	mux.HandleFunc("POST /users/follow/{id}", app.FollowUserHandler)
	// mux.HandleFunc("DELETE /users/{id}/follow", app.UnfollowUserHandler)
	mux.HandleFunc("GET	/users/followers/{id}", app.ListFollowersHandler)
	mux.HandleFunc("GET	/users/following/{id}", app.ListFollowingHandler)

	mux.HandleFunc("POST /messages", app.CreateMessageHandler)
	// mux.HandleFunc("GET /feed", app.feed)
	mux.HandleFunc("GET /messages/{id}", app.GetMessageHandler)
	mux.HandleFunc("GET /users/messages/{id}", app.ListUserMessagesHandler)
	// mux.HandleFunc("DELETE /messages/{id}", app.DeleteMessageHandler)

	// mux.HandleFunc("GET /timeline", app.GetTimelineHandler)
	// mux.HandleFunc("POST /tweets/{id}/retweet", app.RetweetHandler)
	// mux.HandleFunc("DELETE /tweets/{id}/retweet", app.UndoRetweetHandler)
	// mux.HandleFunc("POST /tweets/{id}/favorite", app.FavoriteTweetHandler)
	// mux.HandleFunc("DELETE /tweets/{id}/favorite", app.UnfavoriteTweetHandler)

	// mux.HandleFunc("GET /users/{id}/notifications", app.ListNotificationsHandler)
	// mux.HandleFunc("GET /users/{id}/stats", app.GetUserStatsHandler)
	mux.HandleFunc("POST /auth/login", app.LoginHandler)
	// mux.HandleFunc("POST /auth/logout", app.LogoutHandler)
	mux.HandleFunc("POST /auth/register", app.CreateUserHandler)

	return app.recoverPanic(
		app.logRequest(
			enableCORS(secureHeaders(mux)),
		))
}
