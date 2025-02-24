package main

import (
    "fmt"
    "strings"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)
//USER-AUTH///////////////////////////////////////////////
func showUserActionsScreen(window fyne.Window, app *Application, userID string) {
    // Botones para las acciones disponibles
    followButton := widget.NewButton("Seguir", func() {
        err := app.followUser(userID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Ahora sigues a este usuario", window)
        }
    })

    unfollowButton := widget.NewButton("Dejar de seguir", func() {
        err := app.unfollowUser(userID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Dejaste de seguir a este usuario", window)
        }
    })

    viewProfileButton := widget.NewButton("Ver perfil", func() {
        user, err := app.showProfile(userID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            content := container.NewVBox(
                widget.NewLabel("Nombre de usuario: " + user.UserName),
                widget.NewLabel("Email: " + user.Email),
            )
            dialog.ShowCustom("Perfil de Usuario", "Cerrar", content, window)
        }
    })

    viewMessagesButton := widget.NewButton("Ver mensajes", func() {
        messages, err := app.service.GetUserMessages(userID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            var messageList []fyne.CanvasObject
            for _, message := range messages {
                messageList = append(messageList, widget.NewLabel(message.Content))
            }
            content := container.NewVScroll(container.NewVBox(messageList...))
            dialog.ShowCustom("Mensajes de " + userID, "Cerrar", content, window)
        }
    })

    backButton := widget.NewButton("Volver a la lista de usuarios", func() {
        showUserListScreen(window, app)
    })

    // Contenedor principal
    content := container.NewVBox(
        widget.NewLabel("Acciones para el usuario: " + userID),
        followButton,
        unfollowButton,
        viewProfileButton,
        viewMessagesButton,
        backButton,
    )

    window.SetContent(content)
}

func showLoginScreen(window fyne.Window, app *Application) {
    // Campos de entrada para login
    usernameEntry := widget.NewEntry()
    passwordEntry := widget.NewPasswordEntry()

    // Botones
    loginButton := widget.NewButton("Iniciar Sesión", func() {
        username := usernameEntry.Text
        password := passwordEntry.Text

        // Lógica de autenticación
        err := app.loginComponent(username, password)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            showMainMenu(window, app) // Mostrar el menú principal
        }
    })

    signUpButton := widget.NewButton("Registrarse", func() {
        showSignUpScreen(window, app) // Mostrar la pantalla de registro
    })

    // Contenedor principal
    content := container.NewVBox(
        widget.NewLabel("Bienvenido a la Red Social"),
        widget.NewLabel("Nombre de usuario:"),
        usernameEntry,
        widget.NewLabel("Contraseña:"),
        passwordEntry,
        loginButton,
        signUpButton,
    )

    window.SetContent(content)
}

func showSignUpScreen(window fyne.Window, app *Application) {
    // Campos de entrada para registro
    usernameEntry := widget.NewEntry()
    passwordEntry := widget.NewPasswordEntry()
    emailEntry := widget.NewEntry()

    // Botón de registro
    signUpButton := widget.NewButton("Registrarse", func() {
        username := usernameEntry.Text
        password := passwordEntry.Text
        email := emailEntry.Text

        // Lógica de registro
        err := app.signUpComponent(username, password, email)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            showLoginScreen(window, app) // Volver a la pantalla de login
        }
    })

    // Contenedor principal
    content := container.NewVBox(
        widget.NewLabel("Registro de Usuario"),
        widget.NewLabel("Nombre de usuario:"),
        usernameEntry,
        widget.NewLabel("Contraseña:"),
        passwordEntry,
        widget.NewLabel("Correo electrónico:"),
        emailEntry,
        signUpButton,
    )

    window.SetContent(content)
}

// func showMainMenu(window fyne.Window, app *Application) {
//     // Botones para las funcionalidades
//     buttons := []*widget.Button{
//         widget.NewButton("Ver perfil de usuario", func() { app.showProfile() }),
//         widget.NewButton("Seguir usuario", func() { app.followUser() }),
//         widget.NewButton("Dejar de seguir usuario", func() { app.unfollowUser() }),
//         widget.NewButton("Listar seguidores", func() { app.listFollowers() }),
//         widget.NewButton("Listar seguidos", func() { app.listFollowing() }),
//         widget.NewButton("Crear mensaje", func() { app.createMessageComponent() }),
//         widget.NewButton("Obtener mensaje", func() { app.getMessage() }),
//         widget.NewButton("Eliminar mensaje", func() { app.deleteMessage() }),
//         widget.NewButton("Ver mi perfil", func() { app.showMyProfile() }),
//         widget.NewButton("Actualizar mi perfil", func() { app.updateProfile() }),
//         widget.NewButton("Eliminar mi cuenta", func() { app.deleteUser() }),
//         widget.NewButton("Cerrar sesión", func() {
//             app.token = "" // Cerrar sesión
//             showLoginScreen(window, app) // Volver a la pantalla de login
//         }),
//     }

//     // Contenedor principal
//     content := container.NewVBox(buttons...)
//     window.SetContent(content)
// }

func showMainMenu(window fyne.Window, app *Application) {
    buttons := []*widget.Button{
        widget.NewButton("Ver lista de usuarios", func() {
            showUserListScreen(window, app)
        }),
        widget.NewButton("Crear mensaje", func() {
            showCreateMessageScreen(window, app)
        }),
        widget.NewButton("Ver mi perfil", func() {
            showMyProfileScreen(window, app)
        }),
        widget.NewButton("Cerrar sesión", func() {
            app.token = "" // Cerrar sesión
            showLoginScreen(window, app) // Volver a la pantalla de login
        }),
    }

    var canvasObjects []fyne.CanvasObject
    for _, button := range buttons {
        canvasObjects = append(canvasObjects, button)
    }

    content := container.NewVBox(canvasObjects...)
    window.SetContent(content)
}

func showMyProfileScreen(window fyne.Window, app *Application) {
    user, err := app.showMyProfile()
    if err != nil {
        dialog.ShowError(err, window)
        return
    }

    content := container.NewVBox(
        widget.NewLabel("Mi Perfil"),
        widget.NewLabel("Nombre de usuario: " + user.UserName),
        widget.NewLabel("Email: " + user.Email),
    )

    window.SetContent(content)
}

func showUpdateProfileScreen(window fyne.Window, app *Application) {
    usernameEntry := widget.NewEntry()
    emailEntry := widget.NewEntry()

    updateButton := widget.NewButton("Actualizar", func() {
        username := usernameEntry.Text
        email := emailEntry.Text

        err := app.updateProfile(username, email)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Perfil actualizado correctamente", window)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Actualizar Perfil"),
        widget.NewLabel("Nombre de usuario:"),
        usernameEntry,
        widget.NewLabel("Email:"),
        emailEntry,
        updateButton,
    )

    window.SetContent(content)
}

func showProfileScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    viewButton := widget.NewButton("Ver Perfil", func() {
        userID := idEntry.Text
        user, err := app.showProfile(userID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            content := container.NewVBox(
                widget.NewLabel("Nombre de usuario: " + user.UserName),
                widget.NewLabel("Email: " + user.Email),
            )
            dialog.ShowCustom("Perfil de Usuario", "Cerrar", content, window)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Ver Perfil de Usuario"),
        widget.NewLabel("ID de usuario:"),
        idEntry,
        viewButton,
    )

    window.SetContent(content)
}

func showDeleteAccountScreen(window fyne.Window, app *Application) {
    confirmButton := widget.NewButton("Eliminar Cuenta", func() {
        err := app.deleteUser()
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Cuenta eliminada correctamente", window)
            showLoginScreen(window, app) // Volver a la pantalla de login
        }
    })

    content := container.NewVBox(
        widget.NewLabel("¿Estás seguro de que deseas eliminar tu cuenta?"),
        confirmButton,
    )

    window.SetContent(content)
}

func showUserListScreen(window fyne.Window, app *Application) {
    // Obtener la lista de usuarios usando ListUsers
    users, err := app.service.ListUsers()
    if err != nil {
        dialog.ShowError(fmt.Errorf("no se pudo obtener la lista de usuarios: %v", err), window)
        return
    }

    // Crear un contenedor para la lista de usuarios
    userList := container.NewVBox()

    // Mostrar cada usuario en la lista
    for _, user := range users {
        user := user // Crear una copia local para usar en el closure
        userButton := widget.NewButton(user.UserName, func() {
            showUserActionsScreen(window, app, user.UserID)
        })
        userList.Add(userButton)
    }

    // Botón para volver al menú principal
    backButton := widget.NewButton("Volver al menú principal", func() {
        showMainMenu(window, app)
    })

    // Contenedor principal
    content := container.NewVScroll(container.NewVBox(
        widget.NewLabel("Lista de Usuarios"),
        userList,
        backButton,
    ))

    window.SetContent(content)
}

//MESSAGE/////////////////////////////////////////////////////////////////////////
// Pantalla para crear un mensaje
func showCreateMessageScreen(window fyne.Window, app *Application) {
    contentEntry := widget.NewEntry()

    createButton := widget.NewButton("Publicar", func() {
        content := strings.TrimSpace(contentEntry.Text)
        if content == "" {
            widget.NewLabel("El mensaje no puede estar vacío")
            return
        }

        _, err := app.createMessageComponent(content)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Mensaje publicado correctamente", window)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Crear Mensaje"),
        widget.NewLabel("Contenido:"),
        contentEntry,
        createButton,
    )

    window.SetContent(content)
}

func showGetMessageScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    getButton := widget.NewButton("Obtener Mensaje", func() {
        messageID := idEntry.Text
        message, err := app.getMessage(messageID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            content := container.NewVBox(
                widget.NewLabel("ID del mensaje: " + message.MessageID),
                widget.NewLabel("Contenido: " + message.Content),
                widget.NewLabel("Creado el: " + message.CreatedAt.Format("2006-01-02 15:04:05")),
                widget.NewLabel("Última actualización: " + message.UpdatedAt.Format("2006-01-02 15:04:05")),
            )
            dialog.ShowCustom("Mensaje", "Cerrar", content, window)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Obtener Mensaje"),
        widget.NewLabel("ID del mensaje:"),
        idEntry,
        getButton,
    )

    window.SetContent(content)
}

func showDeleteMessageScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    deleteButton := widget.NewButton("Eliminar Mensaje", func() {
        messageID := idEntry.Text
        err := app.deleteMessage(messageID)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Mensaje eliminado correctamente", window)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Eliminar Mensaje"),
        widget.NewLabel("ID del mensaje:"),
        idEntry,
        deleteButton,
    )

    window.SetContent(content)
}

//INTERACTIONS/////////////////////////////////////////////////////////
// Pantalla para seguir a un usuario
func showFollowUserScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    followButton := widget.NewButton("Seguir", func() {
        id := idEntry.Text
        err := app.followUser(id)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            widget.NewLabel("Ahora sigues a este usuario")
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Seguir Usuario"),
        widget.NewLabel("ID de usuario:"),
        idEntry,
        followButton,
    )

    window.SetContent(content)
}

// Pantalla para dejar de seguir a un usuario
func showUnfollowUserScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    unfollowButton := widget.NewButton("Dejar de seguir", func() {
        id := idEntry.Text
        err := app.unfollowUser(id)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            widget.NewLabel("Dejaste de seguir a este usuario")
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Dejar de Seguir Usuario"),
        widget.NewLabel("ID de usuario:"),
        idEntry,
        unfollowButton,
    )

    window.SetContent(content)
}

// Pantalla para listar seguidores
func showFollowersScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    listButton := widget.NewButton("Listar", func() {
        id := idEntry.Text
        users, err := app.listFollowers(id)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            displayUsers(users)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Listar Seguidores"),
        widget.NewLabel("ID de usuario:"),
        idEntry,
        listButton,
    )

    window.SetContent(content)
}

// Pantalla para listar seguidos
func showFollowingScreen(window fyne.Window, app *Application) {
    idEntry := widget.NewEntry()

    listButton := widget.NewButton("Listar", func() {
        id := idEntry.Text
        users, err := app.listFollowing(id)
        if err != nil {
            widget.NewLabel("Error: " + err.Error())
        } else {
            displayUsers(users)
        }
    })

    content := container.NewVBox(
        widget.NewLabel("Listar Seguidos"),
        widget.NewLabel("ID de usuario:"),
        idEntry,
        listButton,
    )

    window.SetContent(content)
}