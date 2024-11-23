package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.feed)
	mux.HandleFunc("/post/create", app.createPost)

	return mux
}
