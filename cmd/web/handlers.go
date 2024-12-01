package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"distnet/internal/models"
)

func (app *application) feed(w http.ResponseWriter, r *http.Request) { // creo que esto hay que borrarlo
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Hello from Distnet"))
}

func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if payload.UserName == "" || payload.Email == "" || payload.Password == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "missing required fields")
		return
	}

	id, err := app.models.User.Create(payload.UserName, payload.Email, payload.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"user_id": id})

	// w.Write([]byte(fmt.Sprintf("created user %d", id)))
	// http.Redirect(w, r, fmt.Sprintf("/user/%d", id), http.StatusSeeOther)
}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	fmt.Println(id)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.models.User.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		app.serverError(w, err)
	}

	// fmt.Fprintf(w, "%+v", user)

	//w.Write([]byte(fmt.Sprintf("get user %d", user)))
}

func (app *application) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }
	users, err := app.models.User.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// for _, user := range users {
	// 	fmt.Fprintf(w, "%+v\n", user)
	// }
	// w.Write([]byte("Getting users"))
}

func (app *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		app.badRequestResponse(w, r, err)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user.UserID = id
	err = app.models.User.Update(r.Context(), &user)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Or http.StatusNoContent if you don't want to send a response body.
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})

	//w.Write([]byte("Getting users"))
}

func (app *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.User.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK) // or http.StatusNoContent
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

	//w.Write([]byte("Getting users"))
}

/////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var newMessage struct {
	 	UserID  int    `json:"user_id"`
	 	Content string `json:"content"`
	}
	err := decoder.Decode(&newMessage)
	if err != nil {
	 	app.badRequestResponse(w, r, fmt.Errorf("error al decodificar el JSON: %w", err))
	 	return
	}
	defer r.Body.Close()

	if newMessage.UserID == 0 || newMessage.Content == "" {
	 	app.badRequestResponse(w, r, errors.New("user_id y content son requeridos"))
	 	return
	}
	
	messageID, err := app.models.Message.Create(newMessage.UserID, newMessage.Content)
	if err != nil {
	 	if errors.Is(err, models.ErrNoRecord) {
	 	 	app.badRequestResponse(w, r, err)
	 	} else {
	  	app.serverError(w, err)
	 	}
	 	return
	}
	
	createdMessage := models.Message{
	 	MessageID: messageID,
	 	UserID:    newMessage.UserID,
	 	Content:   newMessage.Content,
	 	CreatedAt: time.Now(),
	 	UpdatedAt: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdMessage)
	if err != nil {
	 	app.serverError(w, models.NewErrDatabaseOperationFailed(err))
	}
	   
	//w.Write([]byte("Getting users"))
}
func (app *application) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	
	messageID, err := strconv.Atoi(r.URL.Path[len("/messages/"):])
	if err != nil {
	 	app.badRequestResponse(w, r, fmt.Errorf("ID de mensaje inválido: %w", err))
	 return
	}
   
	message, err := app.models.Message.Get(messageID)
	if err != nil {
	 	if errors.Is(err, models.ErrNoRecord) {
	  	app.notFound(w)
	 	} else {
	  		app.serverError(w, err)
	 	}
	 	return
	}
   
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
	 	app.serverError(w, err)
	}
	
	//w.Write([]byte("Getting users"))
}

func (app *application) ListUserMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) GetTimelineHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

//////////////////////////////////////////////////////////////////////////////////////////////
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

/////////////////////////////////////////////////////////////////////////////////////////////
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
func (app *application) GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

//////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListReceivedMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListSentMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

/////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "missing required fields")
		return
	}

	user, err := app.models.User.Authenticate(payload.Email, payload.Password)
	if err != nil {
		app.invalidCredentialsResponse(w, r) //arreglar esto con el error correspondiente
		return
	}

	token, err := GenerateJWT(user.Username)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": user}, nil)

}

// func (app *application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }

func (app *application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if payload.UserName == "" || payload.Email == "" || payload.Password == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "missing required fields")
		return
	}

	id, err := app.models.User.Create(payload.UserName, payload.Email, payload.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"user_id": id})

	// w.Write([]byte(fmt.Sprintf("created user %d", id)))
	// http.Redirect(w, r, fmt.Sprintf("/user/%d", id), http.StatusSeeOther)
}

