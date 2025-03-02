package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/dto"
	"github.com/NinaSayers/Distributed-Social-Network/server/internal/models"
	"github.com/jbenet/go-base58"
	"golang.org/x/crypto/bcrypt"
)

// func (app *application) feed(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting feed"))
// }

func (app *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var payload dto.CreateUserDTO

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	if payload.UserName == "" || payload.Email == "" || payload.Password == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "missing required fields")
		return
	}

	app.infoLog.Printf("Creating user %s \n", payload.Email)

	enc, err := json.Marshal(payload)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}

	// id, err := app.models.User.Create(payload.UserName, payload.Email, payload.Password)
	_, err = app.peer.Store("user", &enc)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"user_id": 0})

	// w.Write([]byte(fmt.Sprintf("created user %d", id)))
	// http.Redirect(w, r, fmt.Sprintf("/user/%d", id), http.StatusSeeOther)
}

// ///////////////////////////////////////////////////////////////////////////////////////////
func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var payload dto.LoginDTO

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "missing required fields")
		return
	}

	app.infoLog.Printf("Email: %s, Password: %s", payload.Email, payload.Password)

	// user, err := app.models.User.Authenticate(payload.Email, payload.Password)
	hash := sha256.Sum256([]byte(payload.Email)) // More secure hash
	id := base58.Encode(hash[:])

	app.infoLog.Printf("Email %s encripted %s \n", payload.Email, id)

	userBytes, err := app.peer.GetValue("user", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.infoLog.Printf("Usuario obtenido %s", string(userBytes))

	var user dto.AuthUserDTO
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}
	app.infoLog.Printf("User %s retrived \n", user.UserName)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password))
	if err != nil {
		app.errorLog.Println("Validating User", err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			app.invalidCredentialsResponse(w, r)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.infoLog.Println("Successful authentication user ", user.UserName)

	token, err := GenerateJWT(user.UserName)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": user}, nil)

}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing user id"))
		return
	}

	userBytes, err := app.peer.GetValue("user", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.infoLog.Printf("Usuario obtenido %s", string(userBytes))

	var user dto.AuthUserDTO
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}
	app.infoLog.Printf("User %s retrived \n", user.UserName)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		app.serverError(w, err)
	}
}

// func (app *application) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	// if r.URL.Path != "/" {
// 	// 	app.notFound(w)
// 	// 	return
// 	// }
// 	users, err := app.models.User.List()
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	err = json.NewEncoder(w).Encode(users)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	// for _, user := range users {
// 	// 	fmt.Fprintf(w, "%+v\n", user)
// 	// }
// 	// w.Write([]byte("Getting users"))
// }

// func (app *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil || id <= 0 {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	var user models.User
// 	err = json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	user.UserID = id
// 	err = app.models.User.Update(r.Context(), &user)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.badRequestResponse(w, r, err)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK) // Or http.StatusNoContent if you don't want to send a response body.
// 	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})

// 	//w.Write([]byte("Getting users"))
// }

// func (app *application) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil || id <= 0 {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	err = app.models.User.Delete(r.Context(), id)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK) // or http.StatusNoContent
// 	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

// 	//w.Write([]byte("Getting users"))
// }

// // ///////////////////////////////////////////////////////////////////////////////////////////
func (app *application) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Creating post repingaaaaaaa\n")

	var payload dto.CreatePostDTO
	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorLog.Println(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	app.infoLog.Printf("Creating post %s... \n", payload.UserID)

	defer r.Body.Close()

	if payload.UserID == "" || payload.Content == "" {
		app.badRequestResponse(w, r, errors.New("user_id y content son requeridos"))
		return
	}

	// app.infoLog.Printf("Creating post %s... 2\n", payload.Content[:10])

	enc, err := json.Marshal(payload)
	if err != nil {
		app.serverError(w, err)
		return
	}

	createdMessage, err := app.peer.Store("post", &enc)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// createdMessage := models.Message{
	// 	MessageID: messageID,
	// 	UserID:    payload.UserID,
	// 	Content:   payload.Content,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdMessage)
	if err != nil {
		app.serverError(w, models.NewErrDatabaseOperationFailed(err))
	}

	//w.Write([]byte("Getting users"))
}

func (app *application) GetMessageHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, errors.New("missing message id"))
		return
	}

	messageBytes, err := app.peer.GetValue("post", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.infoLog.Printf("Mensaje obtenido %s", string(messageBytes))

	var post models.Post
	err = json.Unmarshal(messageBytes, &post)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}

	app.infoLog.Printf("Mensaje  %s", post.Content)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		app.serverError(w, err)
	}

	//w.Write([]byte("Getting users"))
}

func (app *application) ListUserMessagesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, fmt.Errorf("ID de usuario inválido"))
		return
	}

	// users, err := app.models.Message.ListByUser(int64(id))
	postsBytes, err := app.peer.GetValue("post:user", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	var posts []models.Post
	err = json.Unmarshal(postsBytes, &posts)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}
	app.infoLog.Printf("Post retrived %v\n", len(posts))

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// func (app *application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil || id <= 0 {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	err = app.models.Message.Delete(int64(id))
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.notFound(w)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}

//		w.WriteHeader(http.StatusOK) // or http.StatusNoContent
//		json.NewEncoder(w).Encode(map[string]string{"message": "Post deleted successfully"})
//	}
//
//	func (app *application) GetTimelineHandler(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Getting users"))
//	}
func (app *application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí podrías agregar lógica adicional, como invalidar el token si es necesario.
	app.writeJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"}, nil)
}

// // ////////////////////////////////////////////////////////////////////////////////////////////
func (app *application) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.FollowUserDTO

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if payload.UserID == "" || payload.FolloweeID == "" {
		app.badRequestResponse(w, r, errors.New("user_id y followee_id son requeridos"))
		return
	}

	enc, err := json.Marshal(payload)
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = app.peer.Store("follow", &enc)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// func (app *application) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
// 	followee_id, err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}
// 	var payload struct {
// 		UserID int `json:"user_id"`
// 	}

// 	err = app.readJSON(w, r, &payload)
// 	if err != nil {
// 		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	defer r.Body.Close()

// 	if payload.UserID == 0 {
// 		app.badRequestResponse(w, r, errors.New("relationship_id is required"))
// 		return
// 	}

//		err = app.models.Relationship.UnfollowUser(payload.UserID, followee_id)
//		if err != nil {
//			if errors.Is(err, models.ErrNoRecord) {
//				app.badRequestResponse(w, r, err)
//			} else {
//				app.serverError(w, err)
//			}
//			return
//			// w.Write([]byte("Getting users"))
//		}
//	}
func (app *application) ListFollowersHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, fmt.Errorf("ID de usuario inválido"))
		return
	}

	// users, err := app.models.Message.ListByUser(int64(id))
	followerUsersBytes, err := app.peer.GetValue("follower:user", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	var followUserIDs []string
	err = json.Unmarshal(followerUsersBytes, &followUserIDs)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}
	app.infoLog.Printf("Followers retrived %v\n", len(followUserIDs))

	followers := []*dto.AuthUserDTO{}
	for _, followId := range followUserIDs {
		uBytes, err := app.peer.GetValue("user", followId)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.badRequestResponse(w, r, err)
			} else {
				app.serverError(w, err)
			}
			continue
		}

		var u dto.AuthUserDTO
		err = json.Unmarshal(uBytes, &u)
		if err != nil {
			app.serverError(w, err) //arreglar esto con el error correspondiente
			continue
		}

		followers = append(followers, &u)
	}

	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Getting users"))
}

func (app *application) ListFollowingHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		app.badRequestResponse(w, r, fmt.Errorf("ID de usuario inválido"))
		return
	}

	// users, err := app.models.Message.ListByUser(int64(id))
	followUsersBytes, err := app.peer.GetValue("follow:user", id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.badRequestResponse(w, r, err)
		} else {
			app.serverError(w, err)
		}
		return
	}

	var followUserIDs []string
	err = json.Unmarshal(followUsersBytes, &followUserIDs)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}
	app.infoLog.Printf("Post retrived %v\n", len(followUserIDs))

	following := []*dto.AuthUserDTO{}
	for _, followId := range followUserIDs {
		uBytes, err := app.peer.GetValue("user", followId)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.badRequestResponse(w, r, err)
			} else {
				app.serverError(w, err)
			}
			continue
		}

		var u dto.AuthUserDTO
		err = json.Unmarshal(uBytes, &u)
		if err != nil {
			app.serverError(w, err) //arreglar esto con el error correspondiente
			continue
		}

		following = append(following, &u)
	}

	err = json.NewEncoder(w).Encode(following)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Getting users"))
}

// // ///////////////////////////////////////////////////////////////////////////////////////////
// func (app *application) RetweetHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload struct {
// 		UserID    int `json:"user_id"`
// 		MessageID int `json:"message_id"`
// 	}

// 	err := app.readJSON(w, r, &payload)
// 	if err != nil {
// 		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if payload.UserID == 0 || payload.MessageID == 0 {
// 		app.badRequestResponse(w, r, errors.New("user_id and message_id are required"))
// 		return
// 	}

// 	err = app.models.Retweet.CreateRetweet(payload.UserID, payload.MessageID)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) || errors.Is(err, models.ErrRelationshipExists) {
// 			app.badRequestResponse(w, r, err)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated) // Indicate successful creation

// 	//w.Write([]byte("Getting users"))
// }
// func (app *application) UndoRetweetHandler(w http.ResponseWriter, r *http.Request) {
// 	messageID, err := strconv.Atoi(r.URL.Query().Get("message_id"))
// 	if err != nil {
// 		app.badRequestResponse(w, r, errors.New("invalid message_id"))
// 		return
// 	}
// 	var payload struct {
// 		UserID int `json:"user_id"`
// 	}

// 	err = app.readJSON(w, r, &payload)
// 	if err != nil {
// 		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if payload.UserID == 0 {
// 		app.badRequestResponse(w, r, errors.New("user_id is required"))
// 		return
// 	}

// 	err = app.models.Retweet.UndoRetweet(payload.UserID, messageID)
// 	if err != nil {
// 		if errors.Is(err, models.ErrNoRecord) {
// 			app.badRequestResponse(w, r, err)
// 		} else {
// 			app.serverError(w, err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	//w.Write([]byte("Getting users"))
// }

// // ///////////////////////////////////////////////////////////////////////////////////////////
// func (app *application) FavoriteTweetHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }
// func (app *application) UnfavoriteTweetHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }
// func (app *application) ListNotificationsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }
// func (app *application) GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }

// // ////////////////////////////////////////////////////////////////////////////////////////////
// func (app *application) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }
// func (app *application) ListReceivedMessagesHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }
// func (app *application) ListSentMessagesHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Getting users"))
// }

// // ///////////////////////////////////////////////////////////////////////////////////////////

// // func (app *application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Write([]byte("Getting users"))
// // }

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

	enc, err := json.Marshal(payload)
	if err != nil {
		app.serverError(w, err) //arreglar esto con el error correspondiente
		return
	}

	// id, err := app.models.User.Create(payload.UserName, payload.Email, payload.Password)
	_, err = app.peer.Store("user", &enc)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]int{"user_id": id})

	// w.Write([]byte(fmt.Sprintf("created user %d", id)))
	// http.Redirect(w, r, fmt.Sprintf("/user/%d", id), http.StatusSeeOther)
}
