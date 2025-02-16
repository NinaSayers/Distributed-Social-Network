package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/models"
)

func (app *application) feed(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting feed"))
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
    // Obtener el ID del usuario desde la URL
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id <= 0 {
        app.badRequestResponse(w, r, err)
        return
    }

    // Eliminar el usuario
    err = app.models.User.Delete(r.Context(), id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serverError(w, err)
        }
        return
    }

    // Respuesta exitosa
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

// ///////////////////////////////////////////////////////////////////////////////////////////
func (app *application) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {

	var payload struct {
		UserID  int    `json:"user_id"`
		Content string `json:"content"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if payload.UserID == 0 || payload.Content == "" {
		app.badRequestResponse(w, r, errors.New("user_id y content son requeridos"))
		return
	}

	messageID, err := app.models.Message.Create(payload.UserID, payload.Content)
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
		UserID:    payload.UserID,
		Content:   payload.Content,
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

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("ID de mensaje inválido: %w", err))
		return
	}

	message, err := app.models.Message.Get(id)
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
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("ID de mensaje inválido: %w", err))
		return
	}

	users, err := app.models.Message.ListByUser(int64(id))
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Message.Delete(int64(id))
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK) // or http.StatusNoContent
	response := map[string]string{
        "message": fmt.Sprintf("El mensaje ha sido eliminado correctamente."),
    }
	json.NewEncoder(w).Encode(response)
}
func (app *application) GetTimelineHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

// ////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FollowerID int `json:"follower_id"`
		FolloweeID int `json:"followee_id"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if payload.FollowerID == 0 || payload.FolloweeID == 0 {
		app.badRequestResponse(w, r, errors.New("follower_id y followee_id son requeridos"))
		return
	}

	err = app.models.Relationship.FollowUser(payload.FollowerID, payload.FolloweeID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) { // Assuming ErrNoRecord is defined in your models package
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
		// w.Write([]byte("Getting users"))
	}
}
func (app *application) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followee_id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	var payload struct {
		UserID int `json:"user_id"`
	}

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if payload.UserID == 0 {
		app.badRequestResponse(w, r, errors.New("relationship_id is required"))
		return
	}

	err = app.models.Relationship.UnfollowUser(payload.UserID, followee_id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
		// w.Write([]byte("Getting users"))
	}
}
func (app *application) ListFollowersHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("id")) //Obtener userID de la ruta.  Asumiendo que la ruta es /followers/{id}
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("ID de usuario inválido: %w", err))
		return
	}

	followers, err := app.models.Relationship.ListFollowers(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Getting users"))
}
func (app *application) ListFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("id")) //Obtener userID de la ruta.  Asumiendo que la ruta es /followers/{id}
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("ID de usuario inválido: %w", err))
		return
	}

	followers, err := app.models.Relationship.ListFollowing(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Getting users"))
}

// ///////////////////////////////////////////////////////////////////////////////////////////
func (app *application) RetweetHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID    int `json:"user_id"`
		MessageID int `json:"message_id"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if payload.UserID == 0 || payload.MessageID == 0 {
		app.badRequestResponse(w, r, errors.New("user_id and message_id are required"))
		return
	}

	err = app.models.Retweet.CreateRetweet(payload.UserID, payload.MessageID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) || errors.Is(err, models.ErrRelationshipExists) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusCreated) // Indicate successful creation

	//w.Write([]byte("Getting users"))
}
func (app *application) UndoRetweetHandler(w http.ResponseWriter, r *http.Request) {
	messageID, err := strconv.Atoi(r.URL.Query().Get("message_id"))
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid message_id"))
		return
	}
	var payload struct {
		UserID int `json:"user_id"`
	}

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if payload.UserID == 0 {
		app.badRequestResponse(w, r, errors.New("user_id is required"))
		return
	}

	err = app.models.Retweet.UndoRetweet(payload.UserID, messageID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Getting users"))
}

// ///////////////////////////////////////////////////////////////////////////////////////////
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

// ////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListReceivedMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}
func (app *application) ListSentMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting users"))
}

// ///////////////////////////////////////////////////////////////////////////////////////////
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

func (app *application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
    // Aquí podrías agregar lógica adicional, como invalidar el token si es necesario.
    app.writeJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"}, nil)
}

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
