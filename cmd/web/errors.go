package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

//	func (app *application) logError(r *http.Request, err error) {
//		app.logger.Println(err)
//	}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) { //no idea de donde usar esto
	app.clientError(w, http.StatusNotFound)
}
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// env := envelope{"error": message}
	err := app.writeJSON(w, status, message, nil)
	if err != nil {
		app.errorLog.Println(r, err)
		w.WriteHeader(status)
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
