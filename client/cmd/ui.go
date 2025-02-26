package main

import (
    "fmt"
    "strings"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/widget"
)

func showInitialMenu(window fyne.Window, app *Application) {
    loginButton := widget.NewButton("Iniciar Sesión", func() {
        showLoginScreen(window, app)
    })

    signUpButton := widget.NewButton("Registrarse", func() {
        showSignUpScreen(window, app)
    })

    exitButton := widget.NewButton("Salir", func() {
        window.Close()
    })

    content := container.NewVBox(
        widget.NewLabel("Bienvenido a la Red Social"),
        loginButton,
        signUpButton,
        exitButton,
    )

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
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
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Usuario registrado correctamente", window)
            showInitialMenu(window, app) // Regresar al menú de inicio
        }
    })

    // Botón para volver al menú de inicio
    backButton := widget.NewButton("Volver al menú de inicio", func() {
        showInitialMenu(window, app) // Regresar al menú de inicio
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
        backButton,
    )

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
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
            dialog.ShowError(err, window)
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

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
}

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
            showInitialMenu(window, app) // Volver al menú de inicio
        }),
    }

    var canvasObjects []fyne.CanvasObject
    for _, button := range buttons {
        canvasObjects = append(canvasObjects, button)
    }

    content := container.NewVBox(canvasObjects...)
    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
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

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
}

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

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
}

func showCreateMessageScreen(window fyne.Window, app *Application) {
    contentEntry := widget.NewEntry()

    createButton := widget.NewButton("Publicar", func() {
        content := strings.TrimSpace(contentEntry.Text)
        if content == "" {
            dialog.ShowError(fmt.Errorf("el mensaje no puede estar vacío"), window)
            return
        }

        _, err := app.createMessageComponent(content)
        if err != nil {
            dialog.ShowError(err, window)
        } else {
            dialog.ShowInformation("Éxito", "Mensaje publicado correctamente", window)
        }
    })

    // Botón para volver al menú principal
    backButton := widget.NewButton("Volver al menú principal", func() {
        showMainMenu(window, app)
    })

    // Contenedor principal
    content := container.NewVBox(
        widget.NewLabel("Crear Mensaje"),
        widget.NewLabel("Contenido:"),
        contentEntry,
        createButton,
        backButton,
    )

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
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

    // Botón para volver al menú principal
    backButton := widget.NewButton("Volver al menú principal", func() {
        showMainMenu(window, app)
    })

    content.Add(backButton)

    window.SetContent(content) // Establecer el contenido de la ventana
    window.Content().Refresh() // Forzar la actualización de la ventana
}

