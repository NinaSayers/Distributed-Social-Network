package main

import (
	"encoding/json"
	//"bufio"
	//"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	//"strings"
)

func (app *Application) createMessageComponent(content string) (*Message, error) {
    message, err := app.service.CreateMessage(app.user.UserID, content)
    if err != nil {
        return nil, err
    }
    return message, nil
}

func (app *Application) getMessage(messageID string) (*Message, error) {
    message, err := app.service.GetMessage(messageID)
    if err != nil {
        return nil, err
    }
    return message, nil
}

// Estructura para decodificar la respuesta JSON
type DeleteMessageResponse struct {
    Message string `json:"message"`
}

func (app *Application) deleteMessage(messageID string) error {
    req, err := http.NewRequest(http.MethodDelete, baseURL+"/messages/"+messageID, nil)
    if err != nil {
        return err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var response DeleteMessageResponse
    err = json.Unmarshal(body, &response)
    if err != nil {
        return err
    }

    return nil
}