package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", app.feed)
	mux.HandleFunc("GET	/users", app.ListUsersHandler)
	mux.HandleFunc("POST /users", app.CreateUserHandler)
	mux.HandleFunc("GET /users/{id}", app.GetUserHandler)
	mux.HandleFunc("PUT /users/{id}", app.UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", app.DeleteUserHandler)
	mux.HandleFunc("POST /users/{id}/follow", app.FollowUserHandler)
	mux.HandleFunc("DELETE /users/{id}/follow", app.UnfollowUserHandler)
	mux.HandleFunc("GET	/users/{id}/followers", app.ListFollowersHandler)
	mux.HandleFunc("GET	/users/{id}/following", app.ListFollowingHandler)
	
	mux.HandleFunc("POST /messages", app.CreateMessageHandler)
	// mux.HandleFunc("POST /tweets", app.CreateTweetHandler)
	mux.HandleFunc("GET /messages/{id}", app.GetMessageHandler)
	mux.HandleFunc("GET /users/{id}/messages", app.ListUserMessagesHandler)
	mux.HandleFunc("GET /timeline", app.GetTimelineHandler)
	mux.HandleFunc("DELETE /tweets/{id}", app.DeleteTweetHandler)
	mux.HandleFunc("POST /tweets/{id}/retweet", app.RetweetHandler)
	mux.HandleFunc("DELETE /tweets/{id}/retweet", app.UndoRetweetHandler)
	mux.HandleFunc("POST /tweets/{id}/favorite", app.FavoriteTweetHandler)
	mux.HandleFunc("DELETE /tweets/{id}/favorite", app.UnfavoriteTweetHandler)

	mux.HandleFunc("GET /users/{id}/notifications", app.ListNotificationsHandler)
	mux.HandleFunc("GET /users/{id}/stats", app.GetUserStatsHandler)

	mux.HandleFunc("POST /messages", app.SendMessageHandler)
	mux.HandleFunc("GET /users/{id}/messages/received", app.ListReceivedMessagesHandler)
	mux.HandleFunc("GET /users/{id}/messages/sent", app.ListSentMessagesHandler)

	mux.HandleFunc("POST /auth/login", app.LoginHandler)
	mux.HandleFunc("POST /auth/logout", app.LogoutHandler)
	mux.HandleFunc("POST /auth/register", app.RegisterUserHandler)

	return app.recoverPanic(
		app.logRequest(
			secureHeaders(mux),
		))
}
