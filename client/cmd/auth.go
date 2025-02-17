package main

import (
	"fmt"
)

func (app *Application) signUpComponent() {
	var username, email, password string
	fmt.Print("Nombre de usuario: ")
	fmt.Scan(&username)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	_, err := app.service.CreateUser(username, email, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Usuario creado con éxito")
}

func (app *Application) loginComponent() {
	var email, password string
	fmt.Print("Correo de usuario: ")
	fmt.Scan(&email)
	fmt.Print("Contraseña: ")
	fmt.Scan(&password)

	client, err := app.service.Login(email, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	app.user = &client.User
	app.token = client.Token

	fmt.Printf("Bienvenido %s \n", app.user.UserName)
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

func (app *Application) deleteUser() {
	var confirm string
	fmt.Print("Esta seguro que desea eliminar su cuenta? (S/n): ")
	fmt.Scan(&confirm)

	if confirm == "S" {
		err := app.service.DeleteUser(app.user.UserID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		app.user = nil
		app.token = ""
	}
}
