package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var baseURL string

func init() {
	baseURL = os.Getenv("SERVER_URL")
	if baseURL == "" {
		baseURL = "http://10.0.11.2:4000"
	}
}

func main() {
	var client *Client = nil
	for {
		if client != nil {
			fmt.Println("Hola %s ! Seleccione una opción: ", client.UserName)
			fmt.Println("1. Listar usuarios")
			fmt.Println("3. Obtener usuario")
			fmt.Println("4. Actualizar usuario")
			fmt.Println("5. Eliminar usuario")
			fmt.Println("6. Seguir usuario")
			fmt.Println("7. Dejar de seguir usuario")
			fmt.Println("8. Listar seguidores")
			fmt.Println("9. Listar seguidos")
			fmt.Println("10. Crear mensaje")
			fmt.Println("11. Obtener mensaje")
			fmt.Println("12. Listar mensajes de usuario")
			fmt.Println("13. Eliminar mensaje")
			fmt.Println("15. Salir")

			var option int
			fmt.Scan(&option)

			switch option {
			case 1:
				listUsers()
			case 2:
				signUp()
			case 3:
				getUser()
			case 4:
				updateUser()
			case 5:
				deleteUser()
			case 6:
				followUser()
			case 7:
				unfollowUser()
			case 8:
				listFollowers()
			case 9:
				listFollowing()
			case 10:
				createMessage()
			case 11:
				getMessage()
			case 12:
				listUserMessages()
			case 13:
				deleteMessage()
			case 14:
				login()
			case 15:
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
				c, err := login()
				if err != nil {
					fmt.Println(err)
					continue
				}
				client = c
			case 2:
				signUp()
			case 3:
				os.Exit(0)
			default:
				fmt.Println("Opción no válida")
			}

		}

	}
}

func listUsers() {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func updateUser() {
	var id, username, email string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)
	fmt.Print("Nuevo nombre de usuario: ")
	fmt.Scan(&username)
	fmt.Print("Nuevo email: ")
	fmt.Scan(&email)

	user := map[string]string{"username": username, "email": email}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, baseURL+"/users/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func deleteUser() {
	var id string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+id, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func followUser() {
	var id string
	fmt.Print("ID de usuario a seguir: ")
	fmt.Scan(&id)

	resp, err := http.Post(baseURL+"/users/"+id+"/follow", "application/json", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func unfollowUser() {
	var id string
	fmt.Print("ID de usuario a dejar de seguir: ")
	fmt.Scan(&id)

	req, err := http.NewRequest(http.MethodDelete, baseURL+"/users/"+id+"/follow", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func createMessage() {
	var content string
	fmt.Print("Contenido del mensaje: ")
	fmt.Scan(&content)

	message := map[string]string{"content": content}
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := http.Post(baseURL+"/messages", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func getMessage() {
	var id string
	fmt.Print("ID del mensaje: ")
	fmt.Scan(&id)

	resp, err := http.Get(baseURL + "/messages/" + id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func listUserMessages() {
	var id string
	fmt.Print("ID de usuario: ")
	fmt.Scan(&id)

	resp, err := http.Get(baseURL + "/users/" + id + "/messages")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}

func deleteMessage() {
	var id string
	fmt.Print("ID del mensaje: ")
	fmt.Scan(&id)

	req, err := http.NewRequest(http.MethodDelete, baseURL+"/messages/"+id, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))
}
