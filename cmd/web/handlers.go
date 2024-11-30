package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"distnet/internal/models"
)

func (app *application) feed(w http.ResponseWriter, r *http.Request) { // creo que esto hay que borrarlo
	if r.URL.Path != "/" {
		app.notFound(w, "Invalid path")
		return
	}

	w.Write([]byte("Hello from Distnet"))
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a post"))
}

func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData struct {
		Username      string `json:"username"`
		Email         string `json:"email"`
		PasswordHash string `json:"password_hash"`
	}

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
	 	app.clientError(w, http.StatusBadRequest, "Invalid request payload")
	 	return
	}

	id, err := app.users.Create(userData.Username, userData.Email, userData.PasswordHash)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"user_id": id})


	// w.Write([]byte(fmt.Sprintf("created user %d", id)))
	// http.Redirect(w, r, fmt.Sprintf("/user/%d", id), http.StatusSeeOther)
}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("user_id"))

	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest, "Invalid user ID parameter")
		return
	}

	user, err := app.users.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound, "User not found")
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
	if r.URL.Path != "/" {
		app.notFound(w, "Invalid path at ListUsersHandler")
		return
	}
	users, err := app.users.List()
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

	userIDString := strings.TrimPrefix(r.URL.Path, "/users/")
	userID, err := strconv.Atoi(userIDString)
	if err != nil || userID <= 0 {
		app.clientError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user.UserID = userID
	err = app.users.Update(r.Context(), &user)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound, "User not found")
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

	userIDString := strings.TrimPrefix(r.URL.Path, "/users/")
 	userID, err := strconv.Atoi(userIDString)
 	if err != nil || userID <= 0 {
 	 	app.clientError(w, http.StatusBadRequest, "Invalid user ID")
 	 	return
 	}

 	err = app.users.Delete(r.Context(), userID)
 	if err != nil {
 	 	if errors.Is(err, models.ErrNoRecord) {
 	 	 app.clientError(w, http.StatusNotFound, "User not found")
 	 	} else {
 	  		app.serverError(w, err)
 	 	}
 	 return
 	}

 	w.Header().Set("Content-Type", "application/json")
 	w.WriteHeader(http.StatusOK) // or http.StatusNoContent
 	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

	//w.Write([]byte("Getting users"))
}

////////////////////////////////////////////////////////////////////////////////////////////
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
