package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func (app *Application) createMessageComponent() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Contenido del mensaje: ")
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer el mensaje:", err)
		return
	}

	// Eliminar el salto de línea al final del mensaje
	content = strings.TrimSpace(content)

	// Llamar al servicio para publicar el mensaje
	message, err := app.service.CreateMessage(app.user.UserID, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mostrar el mensaje publicado
	displayPost(*message, *app.user)
}

func (app *Application) getMessage() {
	var id string
	fmt.Print("ID del mensaje: ")
	fmt.Scan(&id)

	message, err := app.service.GetMessage(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayPosts([]Message{*message})
}

// Estructura para decodificar la respuesta JSON
type DeleteMessageResponse struct {
    Message string `json:"message"`
}

func deleteMessage() {
    var id string
    fmt.Print("ID del mensaje: ")
    fmt.Scan(&id)

    // Crear la solicitud HTTP DELETE
    req, err := http.NewRequest(http.MethodDelete, baseURL+"/messages/"+id, nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Enviar la solicitud
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Leer la respuesta del servidor
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Decodificar la respuesta JSON
    var response DeleteMessageResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        fmt.Println("Error decodificando JSON:", err)
        return
    }

    // Mostrar solo el mensaje
    fmt.Println(response.Message)
}
