package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.feed)
	mux.HandleFunc("GET /users", app.GetUserHandler)
	mux.HandleFunc("POST /message", app.createPost)
	mux.HandleFunc("POST /users", app.CreateUserHandler)
	// GET	/users/{id}	GetUserHandler
	// PUT	/users/{id}	UpdateUserHandler
	// DELETE	/users/{id}	DeleteUserHandler
	// GET	/users	ListUsersHandler
	// POST	/users/{id}/follow	FollowUserHandler
	// DELETE	/users/{id}/follow	UnfollowUserHandler
	// GET	/users/{id}/followers	ListFollowersHandler
	// GET	/users/{id}/following	ListFollowingHandler
	// POST	/tweets	CreateTweetHandler
	// GET	/tweets/{id}	GetTweetHandler
	// GET	/users/{id}/tweets	ListUserTweetsHandler
	// GET	/timeline	GetTimelineHandler
	// DELETE	/tweets/{id}	DeleteTweetHandler
	// POST	/tweets/{id}/retweet	RetweetHandler
	// DELETE	/tweets/{id}/retweet	UndoRetweetHandler
	// POST	/tweets/{id}/favorite	FavoriteTweetHandler
	// DELETE	/tweets/{id}/favorite	UnfavoriteTweetHandler
	// GET	/users/{id}/notifications	ListNotificationsHandler
	// POST	/messages	SendMessageHandler
	// GET	/users/{id}/messages/received	ListReceivedMessagesHandler
	// GET	/users/{id}/messages/sent	ListSentMessagesHandler
	// POST	/auth/login	LoginHandler
	// POST	/auth/logout	LogoutHandler
	// POST	/auth/register	RegisterUserHandler
	// GET	/users/{id}/stats	GetUserStatsHandler
	return app.recoverPanic(
		app.logRequest(
			secureHeaders(mux),
		))
}
