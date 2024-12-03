package main

import (
	"fmt"
	"os"
)

var baseURL string

func init() {
	baseURL = os.Getenv("SERVER_URL")
	if baseURL == "" {
		baseURL = "http://10.0.11.2:4000"
	}
}

type Application struct {
	service *Service
	token   string
	user    *User
}

func main() {
	app := Application{service: NewService()}
	for {
		if app.token != "" {
			fmt.Printf("Hola %s! \n", app.user.UserName)
			fmt.Println("Unidos recientemente:")
			app.listUsers()

			fmt.Println("Seleccione una opción:")
			fmt.Println("1. Ver perfil de usuario")

			fmt.Println("2. Seguir usuario")
			fmt.Println("3. Dejar de seguir usuario")

			fmt.Println("4. Listar seguidores")
			fmt.Println("5. Listar seguidos")

			fmt.Println("6. Crear mensaje")
			fmt.Println("7. Obtener mensaje")
			fmt.Println("8. Eliminar mensaje")

			fmt.Println("9. Ver mi perfil")
			fmt.Println("10. Actualizar mi perfil")
			fmt.Println("11. Eliminar mi cuenta")

			fmt.Println("0. Salir")

			var option int
			fmt.Scan(&option)

			switch option {
			case 1:
				app.showProfile()
			case 3:
				// updateUser()
			case 4:

				//
			case 5:
				// followUser(client)
			case 6:
				// unfollowUser(client)
				app.createMessageComponent()

			case 7:
				app.getMessage()
				// listFollowers()
			case 8:
				// listFollowing()
				deleteMessage()

			case 9:
				app.showMyProfile()
			case 10:
				app.updateProfile()
			case 11:
				app.deleteUser()

			case 0:
				os.Exit(0)
			default:
				fmt.Println("Opción no válida")
			}
		} else {
			fmt.Println("Seleccione una opción:")
			fmt.Println("1. Login")
			fmt.Println("2. Crear usuario")
			fmt.Println("3. Salir")

			var option int
			fmt.Scan(&option)

			switch option {
			case 1:
				app.loginComponent()
			case 2:
				app.signUpComponent()
			case 3:
				os.Exit(0)
			default:
				fmt.Println("Opción no válida")
			}

		}

	}
}

func selectMessage() {
	var number int
	fmt.Print("Ingrese el número del mensaje: ")
	fmt.Scan(&number)

}
