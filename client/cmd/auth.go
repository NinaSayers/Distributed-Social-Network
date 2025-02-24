package main

// import (
// 	"fmt"
// )

func (app *Application) signUpComponent(username, email, password string) error {
    _, err := app.service.CreateUser(username, email, password)
    if err != nil {
        return err
    }
    return nil
}

func (app *Application) loginComponent(email, password string) error {
    // Llamar al servicio para autenticar al usuario
    client, err := app.service.Login(email, password)
    if err != nil {
        return err // Devolver el error para manejarlo en la interfaz gráfica
    }

    // Actualizar el estado de la aplicación con el usuario autenticado
    app.user = &client.User
    app.token = client.Token

    return nil // Devolver nil si no hay errores
}

// func (app *Application) logoutComponent() {
//     // Llamar al servicio de Log Out
//     err := app.service.Logout()
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     // Limpiar el estado del usuario
//     app.user = nil
//     app.token = ""

//     fmt.Println("Sesión cerrada correctamente.")
//     //app.showInitialMenu() // Redirigir al menú inicial (para esto hay q modificar la func main del client)
// }

func (app *Application) deleteUser() error {
    err := app.service.DeleteUser(app.user.UserID)
    if err != nil {
        return err
    }

    app.user = nil
    app.token = ""
    return nil
}
